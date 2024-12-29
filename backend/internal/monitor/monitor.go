package monitor

import (
    "os/exec"
    "strconv"
    "strings"
    "time"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
)

type SystemMetrics struct {
    CPU       float64   `json:"cpu"`
    CPUTemp   float64   `json:"cpu_temp,omitempty"`
    Memory    Memory    `json:"memory"`
    Disk      Disk      `json:"disk"`
    Timestamp int64     `json:"timestamp"`
}

type Memory struct {
    Total     uint64  `json:"total"`
    Used      uint64  `json:"used"`
    Available uint64  `json:"available"`
    UsageRate float64 `json:"usage_rate"`
}

type Disk struct {
    Total     uint64  `json:"total"`
    Used      uint64  `json:"used"`
    Available uint64  `json:"available"`
    UsageRate float64 `json:"usage_rate"`
}

// GetSystemMetrics 获取系统指标
func GetSystemMetrics() (*SystemMetrics, error) {
    cpuPercent, err := cpu.Percent(time.Second, false)
    if err != nil {
        return nil, err
    }

    memInfo, err := mem.VirtualMemory()
    if err != nil {
        return nil, err
    }

    diskInfo, err := disk.Usage("/")
    if err != nil {
        return nil, err
    }

    metrics := &SystemMetrics{
        CPU: cpuPercent[0],
        Memory: Memory{
            Total:     memInfo.Total,
            Used:      memInfo.Used,
            Available: memInfo.Available,
            UsageRate: memInfo.UsedPercent,
        },
        Disk: Disk{
            Total:     diskInfo.Total,
            Used:      diskInfo.Used,
            Available: diskInfo.Free,
            UsageRate: diskInfo.UsedPercent,
        },
        Timestamp: time.Now().Unix(),
    }

    // 尝试获取CPU温度
    if temp, err := GetCPUTemperature(); err == nil {
        metrics.CPUTemp = temp
    }

    return metrics, nil
}

// CheckServiceStatus 检查服务状态
func CheckServiceStatus(service string) (bool, error) {
    cmd := exec.Command("systemctl", "is-active", service)
    output, err := cmd.Output()
    if err != nil {
        return false, nil
    }
    return strings.TrimSpace(string(output)) == "active", nil
}

// GetCPUTemperature 获取CPU温度（树莓派特有）
func GetCPUTemperature() (float64, error) {
    cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
    output, err := cmd.Output()
    if err != nil {
        return 0, err
    }
    
    temp, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
    if err != nil {
        return 0, err
    }
    
    return temp / 1000.0, nil // 转换为摄氏度
} 