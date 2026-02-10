# 1. Giai đoạn Build (Thợ hồ)
# Dùng ảnh Go chính chủ để build code
FROM golang:alpine AS builder

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy file quản lý thư viện vào trước (để tận dụng cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ code nguồn vào
COPY . .

# Build ra file thực thi (binary) tên là "main"
RUN go build -o main ./dev/cmd/main.go

# ---------------------------------------------------

# 2. Giai đoạn Run (Shipper)
# Dùng ảnh Alpine siêu nhẹ (chỉ vài MB) để chạy
FROM alpine:latest

WORKDIR /app

# Copy file thực thi từ giai đoạn Build sang
COPY --from=builder /app/main .

# Copy file config và data (nếu cần)
# Lưu ý: Cấu trúc thư mục phải khớp với code ông đọc file
COPY config/ ./config/
# COPY app/repository/data/ ./app/repository/data/ (Nếu ông cần file CSV mẫu)

# Mở cổng 8080
EXPOSE 8080

# Lệnh chạy app
CMD ["./main"]