package naive_ed25519

import (
	"math/big"
	"testing"
)

func TestIsOnCurveCheck(t *testing.T) {
	tests := []struct {
		x  string
		y  string
		on bool
	}{
		{
			"1D607B46AFE6755029CCD30D2C6044D2E94D3F1F0796913BFCB4D7CEB3DC92D",
			"5F2A57658DD54BC1D7774C3B9947DFA5BE3965B6762DF8D5E4D8CE1A3F19BBF2",
			true,
		},
		{
			"C3B830415C61952D5162A3A42BC0BAA54B4AF97C61D39129D65F2E239C31F0F",
			"742EAAC831BBE923DA98B5E55D9404D0CC7B4F154E4959AD2465569DC964BD5F",
			true,
		},
		{
			"C3B830415C61952D5162A3A42BC0BAA54B4AF97C61D39129D65F2E239C31F0F",
			"742EAAC831BBE923DA98B5E55D9404D0CC7B4F154E4959AD2465569DC964BD5E",
			false,
		},
		{
			"3FFDB290D154952CA9B5991B0C02E8F9AC7FD09676D46F4A9E3F1ED65DA8C182",
			"70FE7A52CE54959D80A823EAE1152D1D0DC07D7186648A149321670DA763B4B0",
			true,
		},
	}
	for i, v := range tests {
		x, _ := new(big.Int).SetString(v.x, 16)
		y, _ := new(big.Int).SetString(v.y, 16)
		P := ECPoint{x, y}
		if IsOnCurveCheck(P) == false && v.on == true {
			t.Errorf("Point %d should be on curve", i)
		}
		if IsOnCurveCheck(P) == true && v.on == false {
			t.Errorf("Point %d should not be on curve", i)
		}
	}
}

func TestAddECPoints(t *testing.T) {
	x1, _ := new(big.Int).SetString("DDEA783DA202CB9785925B320CA984985A01CC325C1C004331C457ECCF31290", 16)
	y1, _ := new(big.Int).SetString("6AF30EE5CFA0552320E82A8F03799C70AAD417F3CEE64862A77AF5F43C275C8", 16)
	x2, _ := new(big.Int).SetString("7D4376E95B0C6C6190E776BC99225605ACDA2A6243FEDF645C2A9AC62A55D695", 16)
	y2, _ := new(big.Int).SetString("537713C89AD148C94FDBDEA67F15AB8A8496CFB5F5947ABC4F143DE1E36C9CE0", 16)
	res_x, _ := new(big.Int).SetString("185E45F4D095E33071DBBF10FF664033C6891A88FBAFABD6C0E5AC45EBDEAB9C", 16)
	res_y, _ := new(big.Int).SetString("491649656603D0E0299222AC3A8BB24614AD78B4508492AA3F4A486B564ED6BB", 16)
	a := ECPointGen(x1, y1)
	b := ECPointGen(x2, y2)
	res := ECPointGen(res_x, res_y)

	actual_res := AddECPoints(a, b)
	if actual_res.X.Cmp(res.X) != 0 || actual_res.Y.Cmp(res.Y) != 0 {
		t.Error("Problems with addition")
	}
}

func TestDouble(t *testing.T) {
	x1, _ := new(big.Int).SetString("392A0C1BBE35CA91670A8052E223E5D7E8A9FFA1BC561D851DB94E3407AA43E6", 16)
	y1, _ := new(big.Int).SetString("7F18DB7E9647524F375A9906C459784229F71DEEA47571A6219AF8860022659C", 16)
	P := ECPointGen(x1, y1)
	res_x, _ := new(big.Int).SetString("5F990F5ED7968D918DD7D059F905048C8779B36F12D6A5FEC1AC6E91E3185B1E", 16)
	res_y, _ := new(big.Int).SetString("871EA681A2F7839257D2677A7E9A6092C96DF5F9B128C71F87E58290197E101", 16)
	res := ECPointGen(res_x, res_y)
	actual_res := DoubleECPoints(P)
	if actual_res.X.Cmp(res.X) != 0 || actual_res.Y.Cmp(res.Y) != 0 {
		t.Error("Problems with double")
	}
}

func TestScalarMult(t *testing.T) {
	x1, _ := new(big.Int).SetString("44E7804104BCAFA4249AE3C633281C6D61859C4CB28AD8D17D1568F468131848", 16)
	y1, _ := new(big.Int).SetString("44DC77BFA570A6A0C334284521F5B0090FB502C6E442BCA2F4BD227CD6DFF3EC", 16)
	P := ECPointGen(x1, y1)

	res_x, _ := new(big.Int).SetString("14E3DB699F6965BAA6D67DF3EBF6BF557E17FA51E9C57EF48383FBB275FE75DC", 16)
	res_y, _ := new(big.Int).SetString("12EF57C20714A832539783BEE5A13876B3DDE12EDA4945DBB23BDC9CFBB81607", 16)
	res := ECPointGen(res_x, res_y)

	actual_res := ScalarMult(P, *new(big.Int).SetInt64(64))
	if actual_res.X.Cmp(res.X) != 0 || actual_res.Y.Cmp(res.Y) != 0 {
		t.Error("Problems with addition")
	}
}
