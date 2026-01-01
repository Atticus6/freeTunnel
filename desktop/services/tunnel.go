package services

import (
	"context"
	"sync"

	trycloudflared "github.com/atticus6/freeTunnel/cloudflared"
	"github.com/atticus6/freeTunnel/desktop/database"
	"github.com/atticus6/freeTunnel/desktop/models"
)

type TunnelService struct {
	mu            sync.Mutex
	activeTunnels map[uint]context.CancelFunc
}

func NewTunnelService() *TunnelService {
	return &TunnelService{
		activeTunnels: make(map[uint]context.CancelFunc),
	}
}

func (t *TunnelService) CreateTunnel(name, host string, port int64) (*models.Tunnel, error) {
	tunnel := &models.Tunnel{
		Name: name,
		Port: port,
		Host: host,
	}
	if err := database.DB.Create(tunnel).Error; err != nil {
		return nil, err
	}
	return tunnel, nil
}

func (t *TunnelService) GetAllTunnels() ([]models.Tunnel, error) {
	var tunnels []models.Tunnel
	if err := database.DB.Find(&tunnels).Error; err != nil {
		return nil, err
	}
	return tunnels, nil
}

func (t *TunnelService) DeleteById(id uint) error {
	// 先关闭隧道
	_ = t.CloseTunnel(id)
	return database.DB.Delete(&models.Tunnel{}, id).Error
}

func (t *TunnelService) OpenTunnel(id uint) (err error) {
	var tunnel models.Tunnel

	if err = database.DB.First(&tunnel, id).Error; err != nil {
		return err
	}

	t.mu.Lock()
	// 检查是否已经在运行
	if _, exists := t.activeTunnels[id]; exists {
		t.mu.Unlock()
		return nil
	}
	t.mu.Unlock()

	// 创建可取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 cloudflare tunnel
	quickTunnelUrl, err := trycloudflared.CreateCloudflareTunnel(ctx, int(tunnel.Port), tunnel.Host)
	if err != nil {
		cancel()
		return err
	}

	t.mu.Lock()
	t.activeTunnels[id] = cancel
	t.mu.Unlock()

	// 更新数据库
	tunnel.Active = true
	tunnel.QuickTunnelUrl = quickTunnelUrl

	if err = database.DB.Save(&tunnel).Error; err != nil {
		return err
	}

	return nil
}

func (t *TunnelService) CloseTunnel(id uint) (err error) {
	var tunnel models.Tunnel

	if err = database.DB.First(&tunnel, id).Error; err != nil {
		return err
	}

	t.mu.Lock()
	if cancel, exists := t.activeTunnels[id]; exists {
		cancel()
		delete(t.activeTunnels, id)
	}
	t.mu.Unlock()

	tunnel.Active = false
	tunnel.QuickTunnelUrl = ""

	if err = database.DB.Save(&tunnel).Error; err != nil {
		return err
	}

	return nil
}

func (t *TunnelService) ClearAllQuickTunnelUrls() error {
	// 关闭所有活跃的隧道
	t.mu.Lock()
	for id, cancel := range t.activeTunnels {
		cancel()
		delete(t.activeTunnels, id)
	}
	t.mu.Unlock()

	return database.DB.Model(&models.Tunnel{}).Where("1 = 1").Updates(map[string]interface{}{
		"quick_tunnel_url": "",
		"active":           false,
	}).Error
}
