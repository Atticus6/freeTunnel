package services

import (
	"github.com/atticus6/freeTunnel/desktop/database"
	"github.com/atticus6/freeTunnel/desktop/models"
)

type TunnelService struct{}

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
	return database.DB.Delete(&models.Tunnel{}, id).Error
}

func (t *TunnelService) OpenTunnel(id uint) (err error) {
	var tunnel models.Tunnel

	if err = database.DB.First(&tunnel, id).Error; err != nil {
		return err
	}

	tunnel.Active = true

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

	tunnel.Active = false

	if err = database.DB.Save(&tunnel).Error; err != nil {
		return err
	}

	return nil

}

func (t *TunnelService) ClearAllQuickTunnelUrls() error {
	return database.DB.Model(&models.Tunnel{}).Where("1 = 1").Updates(map[string]interface{}{
		"quick_tunnel_url": "",
		"active":           false,
	}).Error
}
