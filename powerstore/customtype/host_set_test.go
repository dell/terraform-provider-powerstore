package customtype

import (
	"reflect"
	"testing"
)

func TestHostSetType_normalizeStrings(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test with empty input",
			args: args{in: []string{}},
			want: []string{},
		},
		{
			name: "Test with valid CIDRs",
			args: args{in: []string{
				"192.168.1.0/24",
				"10.0.0.0/8",
				"172.16.0.0/12",
			}},
			want: []string{
				"192.168.1.0/24",
				"10.0.0.0/8",
				"172.16.0.0/12",
			},
		},
		{
			name: "Test with valid IP addresses",
			args: args{in: []string{
				"192.168.1.1",
				"10.0.0.1",
				"172.16.0.1",
			}},
			want: []string{
				"192.168.1.1",
				"10.0.0.1",
				"172.16.0.1",
			},
		},
		{
			name: "Test with custom hostnames",
			args: args{in: []string{
				"hostname1",
				"hostname2",
				"hostname3",
			}},
			want: []string{
				"hostname1",
				"hostname2",
				"hostname3",
			},
		},
		{
			name: "Test with invalid CIDR",
			args: args{in: []string{
				"192.168.1.0/33",
			}},
			wantErr: true,
		},
		{
			name: "Test with cidr and its subnet",
			args: args{in: []string{
				"192.168.1.0/24",
				"192.168.1.0/26",
			}},
			want: []string{
				"192.168.1.0/24",
			},
		},
		{
			name: "Test with cidr and its subnet Reversed",
			args: args{in: []string{
				"192.168.1.0/26",
				"192.168.1.0/24",
			}},
			want: []string{
				"192.168.1.0/24",
			},
		},
		{
			name: "Test with an IP which is contained in a CIDR",
			args: args{in: []string{
				"89.208.34.0",
				"89.207.132.170",
				"89.207.1.1/16",
			}},
			want: []string{
				"89.207.0.0/16",
				"89.208.34.0",
			},
		},
		{
			name: "Test with IPv6 IP and CIDR",
			args: args{in: []string{
				"2001:db8:85a3::8a2e:370:7334",
				"2001:db8:85a3::/64",
			}},
			want: []string{
				"2001:db8:85a3::/64",
			},
		},
		{
			name: "Test with IPv6 CIDR which is a subnet of IPv6 CIDR",
			args: args{in: []string{
				"2001:db8:85a3::/48",
				"2001:db8:85a3:1::/64",
			}},
			want: []string{
				"2001:db8:85a3::/48",
			},
		}, {
			name: "Test with normalized IPv6 address",
			args: args{in: []string{
				"2001:db8:85a3::8a2e:370:7334",
			}},
			want: []string{
				"2001:db8:85a3::8a2e:370:7334",
			},
		},
		{
			name: "Test with full IPv6 address",
			args: args{in: []string{
				"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
				"2001:db8:85a3::8a2e:370:7334/128",
			}},
			want: []string{
				"2001:db8:85a3::8a2e:370:7334",
				"2001:db8:85a3::8a2e:370:7334",
			},
		},
		{
			name: "Test with normalized IPv6 CIDR",
			args: args{in: []string{
				"2001:db8:85a3::/64",
			}},
			want: []string{
				"2001:db8:85a3::/64",
			},
		},
		{
			name: "Test with full IPv6 CIDR",
			args: args{in: []string{
				"2001:0db8:85a3:0000:0000:0000:0000:0000/64",
			}},
			want: []string{
				"2001:db8:85a3::/64",
			},
		},
		{
			name: "Test with mixed IPv4, IPv6, and CIDRs with deduplication",
			args: args{in: []string{
				"192.168.1.1",
				"2001:db8:85a3::8a2e:370:7334",
				"192.168.1.0/24",
				"2001:db8:85a3::/64",
				"192.168.1.0",
				"2001:db8:85a3::8a2e:370:7334",
				"hostname1",
			}},
			want: []string{
				"hostname1",
				"192.168.1.0/24",
				"2001:db8:85a3::/64",
			},
		},
		// {
		// 	name: "Another test state",
		// 	args: args{in: []string{
		// 		"192.168.1.0/255.255.255.0",
		// 		"192.168.1.0/255.255.255.192",
		// 		"192.168.1.0/255.255.255.252",
		// 		"2001:db8:85a3::8a2e:370:7334/128",
		// 		"@netgroup1",
		// 		"dellinv.gov",
		// 		"dellinv.gov.in",
		// 		"hostname1",
		// 	}},
		// 	want: []string{
		// 		// "hostname1",
		// 		// "192.168.1.0/255.255.255.0",
		// 		// "192.168.1.0/255.255.255.192",
		// 		// "192.168.1.0/255.255.255.252",
		// 		// "2001:db8:85a3::8a2e:370:7334/128",
		// 	},
		// },
		// {
		// 	name: "Another test plan",
		// 	args: args{in: []string{
		// 		"192.168.1.0/24",
		// 		"192.168.1.0/26",
		// 		"192.168.1.0/30",
		// 		"192.168.1.1/24",
		// 		"192.168.1.2/24",
		// 		"192.168.1.54/255.255.255.0",
		// 		"2001:db8:85a3::8a2e:370:7334",
		// 		"hostname1",
		// 		"@netgroup1",
		// 		"dellinv.gov.in",
		// 		"dellinv.gov",
		// 	}},
		// 	want: []string{
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hst := NewHostSetType()
			got, err := hst.normalizeStrings(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("HostSetType.normalizeStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HostSetType.normalizeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
