ví dụ về unit test trong golang (các lệnh chạy test sẽ thường được chạy trong thư mục chứa file _test
hoặc chỉ định thư mục chứa file test, vd: go test ./calc -v)
1. File _test thường nằm cùng package với file cần test 
2. Quy tắc đặt tên file test là tên file cần test cộng với _test.
ví dụ: tên file cần test là intutils.go, file test sẽ là intutils_test.go
3. Một test case sẽ bắt đầu bằng Test theo sau là tên hàm cần test
ví dụ: func TestIntMinBasic(t *testing.T)
hoặc func TestIntMinTableDriven
4. t.Error* (t.Errorf,...) sẽ báo test failures nhưng vẫn thực hiện test case khác
t.Fail* sẽ báo test failures và dừng không tiếp các test case phía sau nữa
5. Việc viết test case có thể lặp đi lặp lại, vì vậy sẽ sử dụng phong cách  table-driven style.
Đầu vào và đầu ra kỳ vọng sẽ được liệt kê trong một mảng, sau đó sẽ duyệt từng phần tử trong mảng
để thực hiện test logic. t.Run cho phép chạy "subtest"
6. Nếu muốn chỉ định một test case cụ thể thì dùng cờ -run theo sau là tên hàm test
ví dụ:  go test -v -run TestIntMinBasic
7. Nếu muốn chỉ định file thì dùng tên file test theo sau là tên file cần test
vi du: go test -v intutils_test.go intutils.go  