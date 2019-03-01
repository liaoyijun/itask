package server

import (
	"errors"
	"log"
	"strconv"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
)

type Push struct {
	IsDev      bool
	ApnsClient *apns2.Client
}

func (p *Push) getb() []byte {
	//测试证书
	if p.IsDev {
		return []byte{}
	}
	//线上证书
	return []byte{48, 130, 12, 161, 2, 1, 3, 48, 130, 12, 104, 6, 9, 42, 134, 72, 134, 247, 13, 1, 7, 1, 160, 130, 12, 89, 4, 130, 12, 85, 48, 130, 12, 81, 48, 130, 6, 231, 6, 9, 42, 134, 72, 134, 247, 13, 1, 7, 6, 160, 130, 6, 216, 48, 130, 6, 212, 2, 1, 0, 48, 130, 6, 205, 6, 9, 42, 134, 72, 134, 247, 13, 1, 7, 1, 48, 28, 6, 10, 42, 134, 72, 134, 247, 13, 1, 12, 1, 6, 48, 14, 4, 8, 93, 146, 118, 239, 189, 68, 226, 120, 2, 2, 8, 0, 128, 130, 6, 160, 132, 102, 228, 182, 18, 156, 9, 222, 243, 239, 77, 18, 155, 30, 22, 196, 233, 20, 235, 134, 117, 183, 193, 154, 34, 126, 100, 227, 81, 142, 111, 187, 101, 88, 233, 91, 48, 129, 220, 95, 230, 115, 68, 12, 0, 205, 48, 69, 147, 157, 210, 7, 73, 168, 180, 84, 87, 164, 189, 165, 81, 211, 61, 175, 225, 125, 127, 25, 150, 229, 7, 106, 143, 249, 104, 241, 230, 119, 57, 73, 147, 117, 150, 94, 172, 83, 65, 154, 196, 72, 174, 112, 19, 20, 101, 125, 198, 29, 15, 80, 91, 7, 207, 76, 205, 36, 3, 178, 114, 241, 97, 161, 13, 84, 110, 86, 15, 33, 19, 2, 149, 238, 35, 5, 198, 28, 229, 170, 88, 223, 118, 255, 112, 110, 52, 237, 75, 162, 57, 114, 101, 142, 84, 132, 110, 89, 140, 52, 188, 137, 254, 249, 224, 113, 233, 186, 215, 32, 13, 134, 38, 73, 136, 211, 103, 49, 125, 137, 119, 63, 136, 247, 152, 83, 179, 29, 71, 191, 114, 150, 111, 203, 202, 183, 13, 143, 192, 39, 47, 245, 74, 13, 105, 50, 42, 95, 162, 245, 72, 43, 41, 164, 144, 224, 184, 110, 3, 82, 162, 196, 116, 172, 40, 190, 137, 186, 63, 198, 1, 45, 222, 99, 7, 13, 159, 47, 50, 15, 22, 169, 31, 152, 161, 70, 206, 84, 140, 199, 102, 147, 38, 158, 64, 238, 47, 33, 220, 121, 233, 213, 170, 10, 233, 34, 78, 219, 238, 6, 133, 197, 92, 218, 50, 3, 171, 50, 39, 146, 214, 111, 197, 33, 26, 254, 117, 100, 93, 88, 111, 42, 127, 70, 241, 252, 244, 71, 3, 215, 240, 15, 58, 243, 31, 210, 31, 54, 235, 45, 4, 25, 4, 109, 187, 168, 93, 180, 239, 98, 33, 185, 84, 238, 120, 218, 172, 46, 73, 71, 251, 117, 196, 251, 64, 221, 44, 82, 207, 68, 149, 108, 77, 128, 55, 110, 164, 253, 158, 172, 67, 210, 123, 222, 151, 251, 222, 19, 160, 146, 190, 234, 248, 82, 1, 131, 184, 230, 69, 230, 120, 121, 150, 236, 82, 134, 254, 31, 178, 108, 145, 159, 224, 24, 239, 181, 50, 216, 178, 72, 202, 135, 15, 98, 190, 87, 165, 7, 101, 149, 219, 212, 139, 32, 209, 207, 101, 96, 250, 90, 28, 8, 20, 18, 206, 173, 220, 145, 220, 174, 46, 157, 62, 72, 197, 134, 127, 73, 28, 122, 102, 44, 213, 106, 48, 60, 90, 70, 180, 78, 106, 39, 133, 177, 3, 202, 205, 47, 233, 17, 58, 211, 251, 78, 143, 89, 211, 125, 222, 132, 51, 233, 26, 40, 159, 124, 162, 144, 81, 96, 15, 248, 167, 77, 232, 18, 51, 13, 220, 141, 98, 102, 237, 38, 74, 255, 175, 208, 218, 63, 208, 53, 133, 98, 33, 101, 2, 196, 68, 121, 40, 198, 94, 139, 93, 65, 158, 194, 18, 211, 107, 236, 151, 179, 179, 244, 205, 135, 18, 170, 203, 141, 71, 112, 158, 185, 37, 38, 58, 248, 220, 238, 231, 91, 249, 190, 50, 30, 168, 130, 134, 80, 71, 187, 174, 134, 149, 23, 113, 253, 89, 95, 81, 98, 94, 153, 99, 179, 2, 34, 144, 215, 27, 59, 21, 80, 86, 246, 50, 110, 251, 43, 207, 38, 188, 120, 27, 187, 154, 50, 243, 77, 41, 195, 222, 148, 33, 4, 80, 10, 181, 24, 14, 153, 218, 132, 205, 185, 237, 237, 191, 124, 222, 27, 246, 168, 203, 155, 151, 229, 188, 5, 151, 250, 97, 211, 98, 186, 206, 198, 103, 131, 147, 216, 9, 1, 165, 34, 74, 100, 87, 12, 210, 128, 90, 36, 134, 191, 57, 111, 179, 32, 83, 221, 89, 246, 200, 229, 166, 161, 238, 104, 56, 145, 242, 213, 34, 68, 2, 140, 240, 110, 98, 254, 183, 214, 13, 8, 185, 43, 63, 145, 21, 4, 39, 179, 77, 37, 230, 103, 224, 83, 48, 148, 217, 73, 131, 26, 72, 26, 50, 231, 11, 144, 165, 210, 107, 128, 20, 7, 255, 66, 199, 98, 130, 253, 236, 117, 63, 106, 204, 38, 60, 1, 234, 150, 94, 88, 240, 254, 232, 18, 212, 244, 229, 85, 143, 57, 99, 14, 26, 10, 128, 119, 247, 81, 173, 1, 249, 50, 176, 24, 73, 190, 8, 116, 190, 58, 169, 124, 33, 218, 99, 93, 251, 18, 69, 115, 77, 126, 44, 20, 109, 66, 49, 105, 131, 57, 144, 129, 30, 192, 127, 76, 127, 120, 46, 121, 236, 220, 17, 160, 199, 249, 122, 89, 196, 220, 166, 186, 208, 112, 230, 176, 104, 15, 188, 49, 74, 70, 168, 167, 232, 165, 32, 132, 149, 248, 39, 4, 192, 43, 245, 138, 25, 243, 189, 15, 155, 249, 197, 2, 209, 98, 153, 226, 80, 30, 207, 58, 229, 145, 39, 153, 176, 36, 189, 207, 124, 69, 227, 191, 76, 35, 38, 22, 229, 208, 158, 87, 175, 26, 2, 190, 9, 84, 51, 247, 90, 33, 50, 120, 55, 155, 130, 49, 33, 182, 128, 106, 73, 177, 46, 55, 76, 41, 170, 171, 216, 13, 117, 66, 227, 77, 145, 174, 66, 32, 168, 187, 3, 62, 5, 73, 89, 173, 247, 220, 172, 168, 124, 130, 84, 253, 216, 84, 91, 39, 52, 91, 29, 174, 93, 233, 172, 68, 40, 110, 155, 7, 226, 185, 47, 112, 182, 0, 71, 41, 138, 177, 44, 159, 255, 247, 75, 92, 103, 134, 254, 71, 124, 111, 139, 44, 114, 74, 66, 11, 54, 121, 53, 24, 110, 239, 47, 45, 102, 2, 225, 110, 222, 222, 34, 171, 41, 212, 11, 144, 222, 180, 210, 77, 57, 148, 59, 21, 248, 69, 27, 133, 64, 107, 62, 173, 43, 206, 57, 106, 193, 33, 140, 142, 84, 164, 0, 138, 41, 96, 108, 28, 68, 15, 131, 226, 247, 202, 105, 207, 142, 177, 149, 46, 214, 93, 43, 238, 44, 147, 168, 156, 225, 83, 72, 131, 224, 194, 140, 142, 254, 112, 208, 36, 148, 54, 166, 225, 207, 125, 152, 210, 87, 52, 200, 102, 200, 9, 23, 232, 252, 120, 201, 206, 135, 161, 237, 208, 182, 67, 78, 11, 89, 0, 239, 28, 245, 210, 119, 42, 102, 177, 165, 36, 150, 19, 28, 25, 4, 209, 254, 242, 17, 221, 4, 147, 195, 21, 81, 254, 143, 209, 206, 7, 8, 50, 85, 165, 181, 64, 49, 30, 28, 95, 151, 175, 173, 53, 33, 58, 87, 102, 71, 245, 235, 215, 0, 18, 38, 192, 109, 217, 20, 110, 157, 202, 56, 46, 83, 213, 14, 125, 9, 196, 57, 100, 131, 163, 189, 253, 151, 173, 210, 18, 200, 255, 159, 140, 1, 165, 124, 4, 205, 111, 254, 209, 142, 186, 24, 200, 77, 179, 78, 38, 12, 7, 207, 147, 245, 104, 219, 96, 147, 222, 85, 211, 158, 123, 34, 76, 169, 139, 179, 171, 78, 46, 182, 87, 38, 114, 212, 161, 197, 115, 38, 70, 86, 220, 96, 164, 155, 238, 174, 208, 150, 220, 238, 243, 50, 98, 151, 200, 195, 137, 117, 53, 238, 24, 181, 93, 66, 107, 222, 231, 8, 56, 13, 11, 205, 121, 97, 136, 217, 164, 169, 13, 243, 255, 197, 229, 22, 209, 25, 91, 179, 121, 155, 1, 20, 97, 60, 211, 47, 93, 118, 234, 157, 160, 88, 248, 62, 133, 86, 213, 108, 203, 236, 28, 176, 167, 138, 40, 214, 42, 159, 150, 255, 157, 236, 73, 154, 176, 204, 122, 242, 132, 231, 206, 23, 10, 17, 98, 67, 193, 19, 236, 206, 102, 26, 168, 134, 91, 144, 58, 124, 71, 113, 29, 106, 253, 68, 128, 122, 193, 64, 68, 200, 108, 6, 110, 68, 27, 218, 16, 131, 1, 151, 91, 109, 167, 116, 72, 163, 166, 11, 107, 252, 108, 85, 188, 210, 54, 190, 152, 91, 4, 249, 122, 79, 110, 119, 101, 247, 36, 63, 87, 161, 34, 76, 89, 101, 92, 115, 180, 228, 190, 134, 173, 20, 154, 215, 118, 118, 69, 27, 144, 148, 115, 227, 13, 206, 187, 66, 11, 135, 254, 180, 240, 244, 48, 206, 234, 5, 100, 156, 104, 4, 187, 142, 152, 89, 180, 9, 223, 53, 46, 177, 249, 158, 110, 54, 118, 233, 8, 95, 54, 34, 139, 163, 19, 210, 60, 242, 10, 62, 207, 97, 104, 43, 239, 75, 36, 100, 18, 81, 166, 210, 46, 24, 67, 40, 98, 161, 164, 212, 65, 231, 75, 241, 229, 210, 182, 14, 41, 64, 118, 16, 125, 166, 109, 96, 166, 129, 67, 198, 163, 35, 16, 94, 228, 47, 228, 62, 61, 197, 209, 47, 137, 8, 132, 12, 6, 118, 111, 123, 218, 225, 179, 31, 192, 234, 180, 80, 169, 162, 155, 181, 219, 50, 31, 90, 154, 104, 53, 131, 186, 117, 42, 118, 32, 92, 206, 224, 124, 174, 53, 212, 144, 1, 225, 63, 34, 28, 133, 81, 70, 137, 107, 11, 6, 131, 16, 2, 51, 229, 210, 100, 132, 221, 225, 218, 251, 75, 201, 152, 53, 199, 70, 38, 144, 61, 142, 244, 42, 172, 111, 109, 152, 89, 141, 130, 157, 52, 22, 41, 62, 31, 136, 213, 80, 179, 180, 10, 87, 1, 200, 7, 69, 123, 123, 17, 226, 38, 127, 52, 230, 246, 41, 169, 13, 126, 167, 44, 109, 231, 187, 38, 69, 19, 138, 1, 98, 115, 222, 85, 87, 143, 6, 247, 91, 44, 43, 164, 186, 150, 186, 26, 61, 78, 48, 189, 109, 24, 222, 163, 139, 47, 62, 132, 185, 4, 164, 210, 150, 209, 28, 224, 178, 131, 127, 31, 229, 248, 178, 233, 61, 15, 225, 158, 39, 209, 182, 129, 176, 207, 208, 71, 45, 76, 34, 252, 24, 107, 24, 255, 0, 251, 74, 99, 190, 115, 121, 250, 18, 161, 86, 12, 224, 249, 181, 244, 106, 255, 113, 10, 134, 104, 171, 212, 105, 17, 153, 203, 153, 89, 74, 61, 244, 104, 49, 95, 38, 83, 192, 23, 13, 224, 26, 158, 190, 48, 130, 5, 98, 6, 9, 42, 134, 72, 134, 247, 13, 1, 7, 1, 160, 130, 5, 83, 4, 130, 5, 79, 48, 130, 5, 75, 48, 130, 5, 71, 6, 11, 42, 134, 72, 134, 247, 13, 1, 12, 10, 1, 2, 160, 130, 4, 238, 48, 130, 4, 234, 48, 28, 6, 10, 42, 134, 72, 134, 247, 13, 1, 12, 1, 3, 48, 14, 4, 8, 222, 162, 53, 233, 247, 79, 87, 60, 2, 2, 8, 0, 4, 130, 4, 200, 12, 42, 175, 159, 255, 215, 154, 30, 124, 183, 157, 149, 64, 43, 111, 153, 164, 247, 25, 70, 179, 67, 202, 18, 16, 88, 32, 25, 82, 234, 30, 52, 66, 117, 70, 70, 133, 72, 218, 166, 233, 2, 146, 138, 156, 119, 140, 204, 171, 214, 201, 237, 37, 117, 244, 10, 154, 18, 162, 90, 229, 203, 200, 58, 8, 195, 179, 106, 152, 232, 13, 50, 60, 151, 212, 247, 237, 142, 245, 129, 168, 67, 154, 166, 167, 166, 169, 139, 108, 217, 1, 221, 5, 139, 161, 70, 167, 154, 250, 55, 11, 234, 65, 104, 11, 64, 237, 166, 89, 254, 208, 68, 175, 176, 117, 8, 146, 187, 113, 199, 61, 7, 225, 65, 255, 244, 148, 8, 12, 0, 215, 143, 70, 150, 37, 217, 196, 221, 115, 196, 78, 167, 203, 254, 83, 31, 53, 98, 53, 184, 124, 154, 132, 103, 61, 165, 27, 119, 142, 154, 171, 112, 123, 77, 23, 127, 85, 151, 82, 249, 14, 3, 59, 241, 151, 162, 209, 247, 148, 231, 219, 236, 76, 95, 248, 169, 195, 77, 23, 138, 101, 236, 134, 157, 201, 123, 251, 125, 178, 212, 81, 66, 185, 207, 241, 181, 201, 247, 99, 90, 219, 246, 124, 228, 222, 44, 230, 67, 241, 77, 48, 232, 137, 64, 116, 213, 217, 138, 76, 4, 112, 171, 219, 30, 203, 111, 122, 26, 188, 36, 91, 205, 155, 10, 84, 130, 124, 225, 5, 75, 13, 203, 194, 192, 252, 6, 87, 187, 236, 210, 174, 39, 152, 81, 137, 31, 236, 197, 141, 53, 66, 183, 93, 192, 183, 136, 196, 206, 218, 205, 148, 189, 74, 248, 254, 56, 224, 52, 33, 140, 224, 242, 87, 127, 242, 208, 218, 29, 59, 119, 177, 231, 246, 89, 142, 34, 92, 28, 71, 10, 184, 28, 129, 65, 85, 239, 7, 26, 194, 92, 172, 167, 158, 62, 212, 8, 24, 111, 193, 53, 149, 68, 19, 48, 177, 227, 47, 150, 201, 248, 28, 247, 46, 134, 209, 30, 89, 136, 145, 163, 172, 139, 98, 109, 100, 148, 111, 253, 162, 110, 216, 119, 69, 103, 252, 166, 172, 30, 23, 169, 86, 51, 93, 113, 115, 173, 60, 167, 25, 15, 150, 116, 229, 97, 236, 55, 90, 80, 84, 50, 87, 247, 162, 75, 27, 2, 248, 88, 214, 128, 180, 109, 3, 155, 234, 127, 185, 20, 197, 181, 36, 201, 64, 136, 88, 29, 138, 32, 254, 187, 202, 171, 12, 29, 180, 213, 36, 10, 183, 140, 115, 27, 39, 25, 254, 139, 245, 251, 20, 121, 121, 157, 50, 68, 237, 183, 113, 189, 226, 152, 125, 21, 74, 110, 4, 31, 101, 219, 61, 145, 104, 152, 174, 127, 137, 176, 229, 46, 145, 187, 29, 53, 160, 40, 142, 79, 187, 182, 23, 7, 143, 218, 242, 60, 145, 65, 223, 26, 189, 137, 122, 108, 227, 115, 153, 223, 160, 30, 192, 198, 193, 13, 150, 132, 82, 99, 216, 252, 9, 176, 194, 180, 4, 224, 167, 162, 241, 191, 124, 47, 0, 54, 92, 245, 117, 233, 96, 247, 118, 104, 122, 194, 210, 31, 215, 40, 89, 200, 111, 4, 70, 18, 239, 122, 154, 184, 176, 104, 51, 249, 52, 144, 242, 39, 156, 228, 153, 75, 90, 15, 156, 2, 171, 105, 141, 137, 50, 246, 56, 172, 231, 173, 36, 189, 145, 247, 144, 167, 43, 25, 79, 18, 84, 218, 213, 150, 12, 61, 219, 53, 82, 141, 95, 121, 154, 234, 113, 19, 43, 70, 159, 145, 231, 82, 23, 106, 66, 31, 169, 52, 129, 185, 1, 90, 28, 215, 206, 23, 75, 13, 60, 90, 75, 169, 234, 130, 216, 125, 99, 26, 12, 185, 177, 111, 254, 39, 198, 198, 246, 231, 108, 77, 187, 146, 35, 105, 150, 222, 78, 174, 109, 80, 108, 237, 230, 99, 61, 191, 117, 35, 77, 62, 11, 144, 149, 198, 216, 204, 230, 99, 95, 64, 29, 166, 229, 84, 240, 4, 143, 238, 117, 135, 198, 221, 57, 144, 211, 92, 14, 159, 93, 196, 126, 185, 102, 216, 63, 82, 40, 105, 42, 56, 31, 226, 53, 227, 253, 237, 52, 34, 221, 137, 185, 74, 115, 98, 75, 72, 190, 24, 97, 138, 145, 69, 109, 222, 227, 236, 102, 162, 234, 176, 11, 141, 36, 45, 4, 138, 242, 161, 138, 88, 151, 215, 105, 86, 210, 20, 9, 200, 187, 116, 246, 50, 242, 71, 116, 54, 154, 217, 173, 241, 135, 111, 92, 89, 73, 80, 215, 63, 206, 151, 146, 206, 229, 69, 95, 243, 112, 28, 175, 110, 233, 247, 247, 240, 121, 155, 254, 155, 156, 180, 32, 115, 72, 98, 161, 36, 97, 154, 183, 235, 246, 91, 229, 238, 41, 145, 81, 35, 94, 42, 221, 88, 98, 144, 54, 191, 36, 49, 153, 132, 82, 206, 63, 140, 28, 91, 7, 11, 204, 187, 105, 233, 5, 222, 196, 234, 178, 72, 45, 99, 79, 91, 184, 168, 207, 178, 51, 211, 133, 155, 45, 63, 244, 253, 13, 109, 13, 241, 185, 156, 55, 1, 241, 122, 77, 83, 168, 77, 49, 1, 176, 57, 187, 108, 195, 47, 131, 44, 117, 142, 28, 83, 232, 200, 162, 244, 169, 79, 84, 51, 90, 52, 108, 1, 115, 35, 7, 244, 188, 168, 250, 220, 43, 39, 6, 50, 105, 230, 241, 129, 193, 200, 157, 153, 52, 53, 136, 11, 39, 123, 196, 81, 185, 178, 134, 178, 100, 217, 210, 145, 197, 135, 173, 84, 173, 250, 126, 49, 33, 67, 10, 52, 151, 136, 57, 231, 103, 204, 160, 201, 169, 49, 194, 164, 142, 58, 9, 196, 26, 192, 155, 152, 18, 194, 219, 73, 3, 149, 57, 20, 151, 237, 168, 107, 254, 125, 228, 96, 23, 13, 198, 125, 192, 105, 240, 239, 59, 204, 197, 90, 238, 140, 194, 187, 153, 173, 213, 81, 15, 177, 247, 74, 236, 10, 14, 65, 92, 225, 163, 251, 136, 208, 133, 88, 154, 172, 235, 150, 62, 3, 170, 85, 250, 167, 59, 211, 17, 69, 104, 10, 252, 189, 229, 208, 235, 174, 48, 128, 99, 167, 15, 174, 43, 199, 202, 217, 173, 52, 76, 218, 29, 200, 156, 40, 48, 206, 88, 221, 214, 128, 37, 103, 217, 217, 78, 146, 141, 227, 72, 161, 151, 25, 144, 197, 210, 209, 199, 82, 183, 244, 89, 0, 168, 46, 191, 241, 101, 47, 224, 40, 217, 21, 178, 253, 59, 217, 95, 157, 29, 162, 141, 174, 163, 31, 73, 81, 200, 28, 51, 201, 241, 99, 179, 166, 237, 191, 236, 239, 204, 101, 31, 49, 105, 187, 4, 118, 77, 84, 191, 178, 26, 36, 105, 100, 82, 77, 17, 153, 199, 87, 97, 162, 192, 169, 113, 15, 116, 82, 38, 88, 164, 81, 101, 178, 29, 111, 7, 113, 253, 10, 155, 192, 65, 207, 133, 41, 252, 207, 99, 15, 130, 77, 80, 245, 174, 62, 119, 183, 70, 210, 161, 108, 112, 169, 85, 238, 105, 29, 218, 175, 169, 223, 3, 0, 111, 113, 16, 169, 253, 68, 40, 178, 146, 78, 231, 108, 170, 199, 190, 25, 140, 30, 216, 148, 82, 250, 170, 228, 208, 49, 158, 138, 49, 70, 48, 31, 6, 9, 42, 134, 72, 134, 247, 13, 1, 9, 20, 49, 18, 30, 16, 0, 66, 0, 97, 0, 114, 0, 107, 0, 80, 0, 117, 0, 115, 0, 104, 48, 35, 6, 9, 42, 134, 72, 134, 247, 13, 1, 9, 21, 49, 22, 4, 20, 216, 42, 101, 87, 80, 79, 226, 67, 142, 172, 103, 98, 93, 64, 116, 235, 96, 215, 38, 138, 48, 48, 48, 33, 48, 9, 6, 5, 43, 14, 3, 2, 26, 5, 0, 4, 20, 240, 56, 223, 122, 184, 16, 179, 36, 250, 137, 169, 128, 41, 219, 246, 25, 32, 43, 44, 245, 4, 8, 96, 172, 19, 59, 107, 202, 208, 109, 2, 1, 1}
}
func NewPush() *Push {
	p := new(Push)
	cert, err := certificate.FromP12Bytes(p.getb(), "bp")
	if err != nil {
		log.Fatalln("cer error")
	}
	p.ApnsClient = apns2.NewClient(cert).Production()
	return p
}

func (p *Push) PostPush(category string, title string, body string, deviceToken string, params map[string]interface{}) error {

	notification := &apns2.Notification{}
	notification.DeviceToken = deviceToken

	payload := payload.NewPayload().Sound("1107").Category("myNotificationCategory")
	badge := params["badge"]
	if badge != nil {
		badgeStr, pass := badge.(string)
		if pass {
			badgeNum, err := strconv.Atoi(badgeStr)
			if err == nil {
				payload = payload.Badge(badgeNum)
			}
		}
	}

	for key, value := range params {
		payload = payload.Custom(key, value)
	}
	if len(title) > 0 {
		payload.AlertTitle(title)
	}
	if len(body) > 0 {
		payload.AlertBody(body)
	}
	notification.Payload = payload.MutableContent()
	notification.Topic = "me.fin.bark"
	res, err := p.ApnsClient.Push(notification)

	if err != nil {
		log.Println("Error:", err)
		return errors.New("与苹果推送服务器传输数据失败")
	}
	log.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	if res.StatusCode == 200 {
		return nil
	} else {
		return errors.New("推送发送失败 " + res.Reason)
	}

}