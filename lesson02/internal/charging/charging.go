// Package charging calculates EV charging costs in kip.
//
// ENG: Prices are integers in kip per kWh; energy comes in Wh from the meter.
//
// LAO: ຄິດໄລ່ຄ່າສາກໄຟເປັນກີບ — ລາຄາເປັນ integer ກີບ/kWh, ພະລັງງານເປັນ Wh.
package charging

// ConnectorType is the type of charging connector.
type ConnectorType int

const (
	Type2 ConnectorType = iota
	CCS2
	CHAdeMO
)

// PricePerKWh returns the price in kip per kWh for c.
//
// ENG: ok is false if c is not a known connector type.
//
// LAO: ສົ່ງຄືນລາຄາ ກີບ/kWh — ok ເປັນ false ຖ້າ connector type ບໍ່ຮູ້ຈັກ.
func PricePerKWh(c ConnectorType) (price int64, ok bool) {
	switch c {
	case Type2:
		return 1500, true
	case CCS2, CHAdeMO:
		return 2500, true
	default:
		return 0, false
	}
}

// Cost returns the price in kip for charging energyWh on connector c,
// rounded up to a whole kip.
//
// ENG: Returns 0 for an unknown connector type or for energyWh <= 0.
//
// LAO: ສົ່ງຄືນລາຄາເປັນກີບ (ປັດຂຶ້ນ) — ສົ່ງຄືນ 0 ຖ້າ connector ບໍ່ຮູ້ຈັກ ຫຼື energyWh <= 0.
func Cost(c ConnectorType, energyWh int64) int64 {
	price, ok := PricePerKWh(c)
	if !ok || energyWh <= 0 {
		return 0
	}
	// ceil(energyWh * price / 1000) using integers only.
	return (energyWh*price + 999) / 1000
}
