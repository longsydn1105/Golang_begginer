import csv
import random

# Cấu hình
FILE_NAME = 'data_test.csv' # Đổi tên file chút cho đỡ nhầm

NUM_ROWS = 10000000 # 1 Triệu dòng



# Dữ liệu chuẩn (Master Data)
VALID_PRODUCTS = {
    'IP15': 'ELEC',
    'MACBOOK': 'ELEC',
    'SHIRT_L': 'FASH',
}
VALID_WAREHOUSES = ['HN', 'HCM', 'DN']
VALID_TYPES = ['IN']

def generate_data():
    data = []
    
    # Header
    header = ['product_sku', 'category_code', 'warehouse_code', 'quantity', 'transaction_type']
    
    print(f"🔄 Đang sinh {NUM_ROWS} dòng dữ liệu sạch...")

    for i in range(NUM_ROWS):
        # 1. Chọn random nhưng CHỈ LẤY trong tập dữ liệu đúng
        sku_key = random.choice(list(VALID_PRODUCTS.keys()))
        sku = sku_key
        category = VALID_PRODUCTS[sku_key]
        
        warehouse = random.choice(VALID_WAREHOUSES)
        
        # 2. Số lượng luôn dương (Từ 1 đến 500)
        quantity = random.randint(1, 500)
        
        # 3. Type luôn chuẩn
        tx_type = random.choice(VALID_TYPES)
        
        # Thêm vào danh sách (Không chèn logic lỗi nào cả)
        data.append([sku, category, warehouse, quantity, tx_type])

    return header, data

def write_csv(header, data):
    print(f"🚀 Đang ghi dữ liệu vào file {FILE_NAME}...")
    with open(FILE_NAME, mode='w', newline='', encoding='utf-8') as file:
        writer = csv.writer(file)
        writer.writerow(header)
        writer.writerows(data)
    print("✅ Xong! File sạch bong sáng bóng nằm ngay bên cạnh.")

if __name__ == "__main__":
    header, data = generate_data()
    write_csv(header, data)