package mdConvert

import (
	"bufio"
	"fmt"
	"os"
)

// DetectImageFormat 检测图片格式
func DetectImageFormat(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 读取前 8 字节用于格式检测
	reader := bufio.NewReader(file)
	header := make([]byte, 8)
	_, err = reader.Read(header)
	if err != nil {
		return "", err
	}

	// JPEG 头: FF D8 FF
	if len(header) >= 3 && header[0] == 0xFF && header[1] == 0xD8 && header[2] == 0xFF {
		return "JPEG", nil
	}

	// PNG 头: 89 50 4E 47 0D 0A 1A 0A
	if len(header) >= 8 &&
		header[0] == 0x89 && header[1] == 0x50 && header[2] == 0x4E &&
		header[3] == 0x47 && header[4] == 0x0D && header[5] == 0x0A &&
		header[6] == 0x1A && header[7] == 0x0A {
		return "PNG", nil
	}

	// GIF 头: 47 49 46 38 37 61 或 47 49 46 38 39 61
	if len(header) >= 6 &&
		((header[0] == 0x47 && header[1] == 0x49 && header[2] == 0x46 &&
			header[3] == 0x38 && header[4] == 0x37 && header[5] == 0x61) ||
			(header[0] == 0x47 && header[1] == 0x49 && header[2] == 0x46 &&
				header[3] == 0x38 && header[4] == 0x39 && header[5] == 0x61)) {
		return "GIF", nil
	}

	// BMP 头: 42 4D
	if len(header) >= 2 && header[0] == 0x42 && header[1] == 0x4D {
		return "BMP", nil
	}

	return "", fmt.Errorf("无法识别的图片格式")
}
