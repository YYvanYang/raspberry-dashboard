package service

import (
    "fmt"
    "os/exec"
    "strings"
    "sync"
    "time"
    "github.com/patrickmn/go-cache"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
)

type SystemService struct {
    services []string
    cache    *cache.Cache
    mu       sync.RWMutex
}

type SystemStatus struct {
    CPU      float64     `json:"cpu"`
    Memory   MemoryInfo  `json:"memory"`
    Disk     DiskInfo    `json:"disk"`
    Uptime   string      `json:"uptime"`
    Services []Service   `json:"services"`
}

type MemoryInfo struct {
    Total     uint64  `json:"total"`
    Used      uint64  `json:"used"`
    Free      uint64  `json:"free"`
    UsageRate float64 `json:"usage_rate"`
}

type DiskInfo struct {
    Total     uint64  `json:"total"`
    Used      uint64  `json:"used"`
    Free      uint64  `json:"free"`
    UsageRate float64 `json:"usage_rate"`
}

type Service struct {
    Name    string `json:"name"`
    Status  string `json:"status"`
    Running bool   `json:"running"`
}

func NewSystemService() *SystemService {
    return &SystemService{
        services: []string{"hysteria2", "wg-quick@wg0", "cloudflared"},
        cache:    cache.New(5*time.Second, 10*time.Second),
    }
}

func (s *SystemService) GetSystemStatus() (*SystemStatus, error) {
    // 尝试从缓存获取
    if cached, found := s.cache.Get("system_status"); found {
        return cached.(*SystemStatus), nil
    }

    s.mu.Lock()
    defer s.mu.Unlock()

    status, err := s.fetchSystemStatus()
    if err != nil {
        return nil, fmt.Errorf("failed to fetch system status: %w", err)
    }

    // 更新缓存
    s.cache.Set("system_status", status, cache.DefaultExpiration)
    return status, nil
}

func (s *SystemService) fetchSystemStatus() (*SystemStatus, error) {
    var status SystemStatus

    // 获取 CPU 使用率
    cpuPercent, err := cpu.Percent(time.Second, false)
    if err != nil {
        return nil, fmt.Errorf("failed to get CPU usage: %w", err)
    }
    status.CPU = cpuPercent[0]

    // 获取内存信息
    memInfo, err := mem.VirtualMemory()
    if err != nil {
        return nil, fmt.Errorf("failed to get memory info: %w", err)
    }
    status.Memory = MemoryInfo{
        Total:     memInfo.Total,
        Used:      memInfo.Used,
        Free:      memInfo.Free,
        UsageRate: memInfo.UsedPercent,
    }

    // 获取磁盘信息
    diskInfo, err := disk.Usage("/")
    if err != nil {
        return nil, fmt.Errorf("failed to get disk info: %w", err)
    }
    status.Disk = DiskInfo{
        Total:     diskInfo.Total,
        Used:      diskInfo.Used,
        Free:      diskInfo.Free,
        UsageRate: diskInfo.UsedPercent,
    }

    // 获取运行时间
    uptime, err := s.getUptime()
    if err != nil {
        return nil, fmt.Errorf("failed to get uptime: %w", err)
    }
    status.Uptime = uptime

    // 获取服务状态
    services, err := s.getServicesStatus()
    if err != nil {
        return nil, fmt.Errorf("failed to get services status: %w", err)
    }
    status.Services = services

    return &status, nil
}

func (s *SystemService) getUptime() (string, error) {
    cmd := exec.Command("uptime", "-p")
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(output)), nil
}

func (s *SystemService) getServicesStatus() ([]Service, error) {
    var services []Service
    for _, name := range s.services {
        status, err := s.getServiceStatus(name)
        if err != nil {
            return nil, fmt.Errorf("failed to get status for service %s: %w", name, err)
        }
        services = append(services, status)
    }
    return services, nil
}

func (s *SystemService) getServiceStatus(name string) (Service, error) {
    cmd := exec.Command("systemctl", "is-active", name)
    output, err := cmd.Output()
    status := strings.TrimSpace(string(output))
    if err != nil {
        status = "inactive"
    }

    return Service{
        Name:    name,
        Status:  status,
        Running: status == "active",
    }, nil
}

func (s *SystemService) StartService(name string) error {
    if !s.isValidService(name) {
        return fmt.Errorf("invalid service: %s", name)
    }
    cmd := exec.Command("systemctl", "start", name)
    return cmd.Run()
}

func (s *SystemService) StopService(name string) error {
    if !s.isValidService(name) {
        return fmt.Errorf("invalid service: %s", name)
    }
    cmd := exec.Command("systemctl", "stop", name)
    return cmd.Run()
}

func (s *SystemService) RestartService(name string) error {
    if !s.isValidService(name) {
        return fmt.Errorf("invalid service: %s", name)
    }
    cmd := exec.Command("systemctl", "restart", name)
    return cmd.Run()
}

func (s *SystemService) isValidService(name string) bool {
    for _, service := range s.services {
        if service == name {
            return true
        }
    }
    return false
} 