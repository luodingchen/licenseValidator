package models

type License struct {
	Config           Config `json:"config" yaml:"-"`
	ConfigJson       string `json:"config_json" yaml:"config_json"`
	LicensePublicKey string `json:"license_public_key" yaml:"license_public_key"`
	LicenseSignature string `json:"license_signature" yaml:"license_signature"`
}

type Config struct {
	StartTime    string         `json:"start_time"`
	Deadline     string         `json:"deadline"`
	ProductList  []Product      `json:"product_list"`
	HardwareList []HardwareInfo `json:"hardware_list"`
}

type Product struct {
	ProductCode        string `json:"product_code"` //字母 -
	ProductName        string `json:"product_name"`
	ProductVersion     string `json:"product_version"` // 1.0 2.0
	ProductStatus      uint   `json:"product_status"`  // 产品关闭时不再售卖 0开启 1关闭
	ProductDescription string `json:"product_description"`
	ProductFuncList    []Func `json:"product_func_list"`
}

type Func struct {
	FunctionCode  string `form:"function_code" json:"function_code"`
	FunctionName  string `form:"function_name" json:"function_name"`   // 可以为中文
	FunctionValue string `form:"function_value" json:"function_value"` // 可以为中文
}
