package utils

type SuccessMessagePrototype struct {
	APIVersion string     `json:"apiVersion"`
	Data       DataObject `json:"data"`
}

type DataObject struct {
	Kind        *string     `json:"kind,omitempty"`
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Item        interface{} `json:"item,omitempty"`
	Items       interface{} `json:"items,omitempty"`
}

func SuccessMessage(data DataObject) SuccessMessagePrototype {
	return SuccessMessagePrototype{APIVersion: "v1", Data: data}
}
