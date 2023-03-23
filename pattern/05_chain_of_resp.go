package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Service interface {
	Execute(*Data)
	SetNext(Service)
}
type Data struct {
	GetSource    bool
	UpdateSource bool
}

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data from device [%s] already get.\n", d.Name)
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s].\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(svc Service) {
	d.Next = svc
}

type UpdateDataService struct {
	Name string
	Next Service
}

func (u *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data in service [%s] is already update.\n", u.Name)
		u.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from service [%s].\n", u.Name)
	data.UpdateSource = true
	u.Next.Execute(data)
}

func (u *UpdateDataService) SetNext(svc Service) {
	u.Next = svc
}

type DataService struct {
	Next Service
}

func (d *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Printf("Data not update.\n")
		return
	}
	fmt.Printf("Data save.\n")
}

func (d *DataService) SetNext(svc Service) {
	d.Next = svc
}

func main() {
	device := &Device{Name: "Device-1"}
	updateSvc := &UpdateDataService{Name: "Update-1"}
	dataSvc := &DataService{}

	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)

	data := &Data{}

	device.Execute(data)
}
