package addr

/*
func TestCheckLocalPort(t *testing.T) {
	ls, err := net.Listen("tcp", "0.0.0.0:12345")
	if err == nil {
		using, _, err := CheckLocalPort("tcp", 12345)
		if err == nil {
			if !using {
				t.Error("CheckLocalPort failed")
			}
		} else {
			t.Error(err)
		}
		ls.C()
	} else {
		t.Error(err)
	}

	using, _, err := CheckLocalPort("tcp", 12345)
	if err == nil {
		if using {
			t.Error("CheckLocalPort failed")
		}
	} else {
		t.Error(err)
	}


	address, err := net.ResolveUDPAddr("udp", "0.0.0.0:12345")
	if err != nil {
		t.Error(err)
		return
	}
	conn, err := net.ListenUDP("udp", address)
	if err == nil {
		using, _, err = CheckLocalPort("udp", 12345)
		if err == nil {
			if !using {
				t.Error("CheckLocalPort failed")
			}
		} else {
			t.Error(err)
		}
		conn.C()
	} else {
		t.Error(err)
	}


	using, _, err = CheckLocalPort("udp", 12345)
	if err == nil {
		if using {
			t.Error("CheckLocalPort failed")
		}
	} else {
		t.Error(err)
	}
}*/
