package charging

import "testing"

// TestCost checks Cost for exact, rounded-up and invalid inputs.
//
// ENG: Each case runs as a subtest, e.g. TestCost/1_Wh_rounds_up.
//
// LAO: ກວດ Cost ຫຼາຍກໍລະນີ — ພໍດີ, ປັດຂຶ້ນ, ແລະ ຄ່າທີ່ບໍ່ຖືກຕ້ອງ.
func TestCost(t *testing.T) {
	tests := []struct {
		name     string
		conn     ConnectorType
		energyWh int64
		want     int64
	}{
		{name: "1 kWh exact", conn: CCS2, energyWh: 1000, want: 2500},
		{name: "1 Wh rounds up", conn: CCS2, energyWh: 1, want: 3},
		{name: "zero energy", conn: CCS2, energyWh: 0, want: 0},
		{name: "Type2 half kip", conn: Type2, energyWh: 1, want: 2},
		{name: "Type2 999 Wh", conn: Type2, energyWh: 999, want: 1499},
		{name: "CHAdeMO 12345 Wh", conn: CHAdeMO, energyWh: 12345, want: 30863},
		{name: "unknown connector", conn: ConnectorType(42), energyWh: 1000, want: 0},
		{name: "negative energy", conn: CCS2, energyWh: -5, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Cost(tc.conn, tc.energyWh)
			if got != tc.want {
				t.Errorf("Cost(%d, %d) = %d, want %d", tc.conn, tc.energyWh, got, tc.want)
			}
		})
	}
}
