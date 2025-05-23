/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// TimeZoneEnum Time zone identifier for applying the time zone to the time_of_day for a snapshot rule, including any DST effects if applicable. Applies only when a time_of_day is specified in the snapshot rule. Defaults to UTC if not specified. Values are:   * Etc__GMT_plus_12   * US__Samoa   * Etc__GMT_plus_11   * America__Atka   * US__Hawaii   * Etc__GMT_plus_10   * Pacific__Marquesas   * US__Alaska   * Pacific__Gambier   * Etc__GMT_plus_9   * PST8PDT   * Pacific__Pitcairn   * US__Pacific   * Etc__GMT_plus_8   * Mexico__BajaSur   * America__Boise   * America__Phoenix   * MST7MDT   * Etc__GMT_plus_7   * CST6CDT   * America__Chicago   * Canada__Saskatchewan   * America__Bahia_Banderas   * Etc__GMT_plus_6   * Chile__EasterIsland   * America__Bogota   * America__New_York   * EST5EDT   * America__Havana   * Etc__GMT_plus_5   * America__Caracas   * America__Cuiaba   * America__Santo_Domingo   * Canada__Atlantic   * America__Asuncion   * Etc__GMT_plus_4   * Canada__Newfoundland   * Chile__Continental   * Brazil__East   * America__Godthab   * America__Miquelon   * America__Buenos_Aires   * Etc__GMT_plus_3   * America__Noronha   * Etc__GMT_plus_2   * America__Scoresbysund   * Atlantic__Cape_Verde   * Etc__GMT_plus_1   * UTC   * Europe__London   * Africa__Casablanca   * Atlantic__Reykjavik   * Antarctica__Troll   * Europe__Paris   * Europe__Sarajevo   * Europe__Belgrade   * Europe__Rome   * Africa__Tunis   * Etc__GMT_minus_1   * Asia__Gaza   * Europe__Bucharest   * Europe__Helsinki   * Asia__Beirut   * Africa__Harare   * Asia__Damascus   * Asia__Amman   * Europe__Tiraspol   * Asia__Jerusalem   * Etc__GMT_minus_2   * Asia__Baghdad   * Africa__Asmera   * Etc__GMT_minus_3   * Asia__Tehran   * Asia__Baku   * Etc__GMT_minus_4   * Asia__Kabul   * Asia__Karachi   * Etc__GMT_minus_5   * Asia__Kolkata   * Asia__Katmandu   * Asia__Almaty   * Etc__GMT_minus_6   * Asia__Rangoon   * Asia__Hovd   * Asia__Bangkok   * Etc__GMT_minus_7   * Asia__Hong_Kong   * Asia__Brunei   * Asia__Singapore   * Etc__GMT_minus_8   * Asia__Pyongyang   * Australia__Eucla   * Asia__Seoul   * Etc__GMT_minus_9   * Australia__Darwin   * Australia__Adelaide   * Australia__Sydney   * Australia__Brisbane   * Asia__Magadan   * Etc__GMT_minus_10   * Australia__Lord_Howe   * Etc__GMT_minus_11   * Asia__Kamchatka   * Pacific__Fiji   * Antarctica__South_Pole   * Etc__GMT_minus_12   * Pacific__Chatham   * Pacific__Tongatapu   * Pacific__Apia   * Etc__GMT_minus_13   * Pacific__Kiritimati   * Etc__GMT_minus_14  Was added in version 2.0.0.0.
type TimeZoneEnum string

// List of TimeZoneEnum
const (
	TIMEZONEENUM_ETC__GMT_PLUS_12        TimeZoneEnum = "Etc__GMT_plus_12"
	TIMEZONEENUM_US__SAMOA               TimeZoneEnum = "US__Samoa"
	TIMEZONEENUM_ETC__GMT_PLUS_11        TimeZoneEnum = "Etc__GMT_plus_11"
	TIMEZONEENUM_AMERICA__ATKA           TimeZoneEnum = "America__Atka"
	TIMEZONEENUM_US__HAWAII              TimeZoneEnum = "US__Hawaii"
	TIMEZONEENUM_ETC__GMT_PLUS_10        TimeZoneEnum = "Etc__GMT_plus_10"
	TIMEZONEENUM_PACIFIC__MARQUESAS      TimeZoneEnum = "Pacific__Marquesas"
	TIMEZONEENUM_US__ALASKA              TimeZoneEnum = "US__Alaska"
	TIMEZONEENUM_PACIFIC__GAMBIER        TimeZoneEnum = "Pacific__Gambier"
	TIMEZONEENUM_ETC__GMT_PLUS_9         TimeZoneEnum = "Etc__GMT_plus_9"
	TIMEZONEENUM_PST8_PDT                TimeZoneEnum = "PST8PDT"
	TIMEZONEENUM_PACIFIC__PITCAIRN       TimeZoneEnum = "Pacific__Pitcairn"
	TIMEZONEENUM_US__PACIFIC             TimeZoneEnum = "US__Pacific"
	TIMEZONEENUM_ETC__GMT_PLUS_8         TimeZoneEnum = "Etc__GMT_plus_8"
	TIMEZONEENUM_MEXICO__BAJA_SUR        TimeZoneEnum = "Mexico__BajaSur"
	TIMEZONEENUM_AMERICA__BOISE          TimeZoneEnum = "America__Boise"
	TIMEZONEENUM_AMERICA__PHOENIX        TimeZoneEnum = "America__Phoenix"
	TIMEZONEENUM_MST7_MDT                TimeZoneEnum = "MST7MDT"
	TIMEZONEENUM_ETC__GMT_PLUS_7         TimeZoneEnum = "Etc__GMT_plus_7"
	TIMEZONEENUM_CST6_CDT                TimeZoneEnum = "CST6CDT"
	TIMEZONEENUM_AMERICA__CHICAGO        TimeZoneEnum = "America__Chicago"
	TIMEZONEENUM_CANADA__SASKATCHEWAN    TimeZoneEnum = "Canada__Saskatchewan"
	TIMEZONEENUM_AMERICA__BAHIA_BANDERAS TimeZoneEnum = "America__Bahia_Banderas"
	TIMEZONEENUM_ETC__GMT_PLUS_6         TimeZoneEnum = "Etc__GMT_plus_6"
	TIMEZONEENUM_CHILE__EASTER_ISLAND    TimeZoneEnum = "Chile__EasterIsland"
	TIMEZONEENUM_AMERICA__BOGOTA         TimeZoneEnum = "America__Bogota"
	TIMEZONEENUM_AMERICA__NEW_YORK       TimeZoneEnum = "America__New_York"
	TIMEZONEENUM_EST5_EDT                TimeZoneEnum = "EST5EDT"
	TIMEZONEENUM_AMERICA__HAVANA         TimeZoneEnum = "America__Havana"
	TIMEZONEENUM_ETC__GMT_PLUS_5         TimeZoneEnum = "Etc__GMT_plus_5"
	TIMEZONEENUM_AMERICA__CARACAS        TimeZoneEnum = "America__Caracas"
	TIMEZONEENUM_AMERICA__CUIABA         TimeZoneEnum = "America__Cuiaba"
	TIMEZONEENUM_AMERICA__SANTO_DOMINGO  TimeZoneEnum = "America__Santo_Domingo"
	TIMEZONEENUM_CANADA__ATLANTIC        TimeZoneEnum = "Canada__Atlantic"
	TIMEZONEENUM_AMERICA__ASUNCION       TimeZoneEnum = "America__Asuncion"
	TIMEZONEENUM_ETC__GMT_PLUS_4         TimeZoneEnum = "Etc__GMT_plus_4"
	TIMEZONEENUM_CANADA__NEWFOUNDLAND    TimeZoneEnum = "Canada__Newfoundland"
	TIMEZONEENUM_CHILE__CONTINENTAL      TimeZoneEnum = "Chile__Continental"
	TIMEZONEENUM_BRAZIL__EAST            TimeZoneEnum = "Brazil__East"
	TIMEZONEENUM_AMERICA__GODTHAB        TimeZoneEnum = "America__Godthab"
	TIMEZONEENUM_AMERICA__MIQUELON       TimeZoneEnum = "America__Miquelon"
	TIMEZONEENUM_AMERICA__BUENOS_AIRES   TimeZoneEnum = "America__Buenos_Aires"
	TIMEZONEENUM_ETC__GMT_PLUS_3         TimeZoneEnum = "Etc__GMT_plus_3"
	TIMEZONEENUM_AMERICA__NORONHA        TimeZoneEnum = "America__Noronha"
	TIMEZONEENUM_ETC__GMT_PLUS_2         TimeZoneEnum = "Etc__GMT_plus_2"
	TIMEZONEENUM_AMERICA__SCORESBYSUND   TimeZoneEnum = "America__Scoresbysund"
	TIMEZONEENUM_ATLANTIC__CAPE_VERDE    TimeZoneEnum = "Atlantic__Cape_Verde"
	TIMEZONEENUM_ETC__GMT_PLUS_1         TimeZoneEnum = "Etc__GMT_plus_1"
	TIMEZONEENUM_UTC                     TimeZoneEnum = "UTC"
	TIMEZONEENUM_EUROPE__LONDON          TimeZoneEnum = "Europe__London"
	TIMEZONEENUM_AFRICA__CASABLANCA      TimeZoneEnum = "Africa__Casablanca"
	TIMEZONEENUM_ATLANTIC__REYKJAVIK     TimeZoneEnum = "Atlantic__Reykjavik"
	TIMEZONEENUM_ANTARCTICA__TROLL       TimeZoneEnum = "Antarctica__Troll"
	TIMEZONEENUM_EUROPE__PARIS           TimeZoneEnum = "Europe__Paris"
	TIMEZONEENUM_EUROPE__SARAJEVO        TimeZoneEnum = "Europe__Sarajevo"
	TIMEZONEENUM_EUROPE__BELGRADE        TimeZoneEnum = "Europe__Belgrade"
	TIMEZONEENUM_EUROPE__ROME            TimeZoneEnum = "Europe__Rome"
	TIMEZONEENUM_AFRICA__TUNIS           TimeZoneEnum = "Africa__Tunis"
	TIMEZONEENUM_ETC__GMT_MINUS_1        TimeZoneEnum = "Etc__GMT_minus_1"
	TIMEZONEENUM_ASIA__GAZA              TimeZoneEnum = "Asia__Gaza"
	TIMEZONEENUM_EUROPE__BUCHAREST       TimeZoneEnum = "Europe__Bucharest"
	TIMEZONEENUM_EUROPE__HELSINKI        TimeZoneEnum = "Europe__Helsinki"
	TIMEZONEENUM_ASIA__BEIRUT            TimeZoneEnum = "Asia__Beirut"
	TIMEZONEENUM_AFRICA__HARARE          TimeZoneEnum = "Africa__Harare"
	TIMEZONEENUM_ASIA__DAMASCUS          TimeZoneEnum = "Asia__Damascus"
	TIMEZONEENUM_ASIA__AMMAN             TimeZoneEnum = "Asia__Amman"
	TIMEZONEENUM_EUROPE__TIRASPOL        TimeZoneEnum = "Europe__Tiraspol"
	TIMEZONEENUM_ASIA__JERUSALEM         TimeZoneEnum = "Asia__Jerusalem"
	TIMEZONEENUM_ETC__GMT_MINUS_2        TimeZoneEnum = "Etc__GMT_minus_2"
	TIMEZONEENUM_ASIA__BAGHDAD           TimeZoneEnum = "Asia__Baghdad"
	TIMEZONEENUM_AFRICA__ASMERA          TimeZoneEnum = "Africa__Asmera"
	TIMEZONEENUM_ETC__GMT_MINUS_3        TimeZoneEnum = "Etc__GMT_minus_3"
	TIMEZONEENUM_ASIA__TEHRAN            TimeZoneEnum = "Asia__Tehran"
	TIMEZONEENUM_ASIA__BAKU              TimeZoneEnum = "Asia__Baku"
	TIMEZONEENUM_ETC__GMT_MINUS_4        TimeZoneEnum = "Etc__GMT_minus_4"
	TIMEZONEENUM_ASIA__KABUL             TimeZoneEnum = "Asia__Kabul"
	TIMEZONEENUM_ASIA__KARACHI           TimeZoneEnum = "Asia__Karachi"
	TIMEZONEENUM_ETC__GMT_MINUS_5        TimeZoneEnum = "Etc__GMT_minus_5"
	TIMEZONEENUM_ASIA__KOLKATA           TimeZoneEnum = "Asia__Kolkata"
	TIMEZONEENUM_ASIA__KATMANDU          TimeZoneEnum = "Asia__Katmandu"
	TIMEZONEENUM_ASIA__ALMATY            TimeZoneEnum = "Asia__Almaty"
	TIMEZONEENUM_ETC__GMT_MINUS_6        TimeZoneEnum = "Etc__GMT_minus_6"
	TIMEZONEENUM_ASIA__RANGOON           TimeZoneEnum = "Asia__Rangoon"
	TIMEZONEENUM_ASIA__HOVD              TimeZoneEnum = "Asia__Hovd"
	TIMEZONEENUM_ASIA__BANGKOK           TimeZoneEnum = "Asia__Bangkok"
	TIMEZONEENUM_ETC__GMT_MINUS_7        TimeZoneEnum = "Etc__GMT_minus_7"
	TIMEZONEENUM_ASIA__HONG_KONG         TimeZoneEnum = "Asia__Hong_Kong"
	TIMEZONEENUM_ASIA__BRUNEI            TimeZoneEnum = "Asia__Brunei"
	TIMEZONEENUM_ASIA__SINGAPORE         TimeZoneEnum = "Asia__Singapore"
	TIMEZONEENUM_ETC__GMT_MINUS_8        TimeZoneEnum = "Etc__GMT_minus_8"
	TIMEZONEENUM_ASIA__PYONGYANG         TimeZoneEnum = "Asia__Pyongyang"
	TIMEZONEENUM_AUSTRALIA__EUCLA        TimeZoneEnum = "Australia__Eucla"
	TIMEZONEENUM_ASIA__SEOUL             TimeZoneEnum = "Asia__Seoul"
	TIMEZONEENUM_ETC__GMT_MINUS_9        TimeZoneEnum = "Etc__GMT_minus_9"
	TIMEZONEENUM_AUSTRALIA__DARWIN       TimeZoneEnum = "Australia__Darwin"
	TIMEZONEENUM_AUSTRALIA__ADELAIDE     TimeZoneEnum = "Australia__Adelaide"
	TIMEZONEENUM_AUSTRALIA__SYDNEY       TimeZoneEnum = "Australia__Sydney"
	TIMEZONEENUM_AUSTRALIA__BRISBANE     TimeZoneEnum = "Australia__Brisbane"
	TIMEZONEENUM_ASIA__MAGADAN           TimeZoneEnum = "Asia__Magadan"
	TIMEZONEENUM_ETC__GMT_MINUS_10       TimeZoneEnum = "Etc__GMT_minus_10"
	TIMEZONEENUM_AUSTRALIA__LORD_HOWE    TimeZoneEnum = "Australia__Lord_Howe"
	TIMEZONEENUM_ETC__GMT_MINUS_11       TimeZoneEnum = "Etc__GMT_minus_11"
	TIMEZONEENUM_ASIA__KAMCHATKA         TimeZoneEnum = "Asia__Kamchatka"
	TIMEZONEENUM_PACIFIC__FIJI           TimeZoneEnum = "Pacific__Fiji"
	TIMEZONEENUM_ANTARCTICA__SOUTH_POLE  TimeZoneEnum = "Antarctica__South_Pole"
	TIMEZONEENUM_ETC__GMT_MINUS_12       TimeZoneEnum = "Etc__GMT_minus_12"
	TIMEZONEENUM_PACIFIC__CHATHAM        TimeZoneEnum = "Pacific__Chatham"
	TIMEZONEENUM_PACIFIC__TONGATAPU      TimeZoneEnum = "Pacific__Tongatapu"
	TIMEZONEENUM_PACIFIC__APIA           TimeZoneEnum = "Pacific__Apia"
	TIMEZONEENUM_ETC__GMT_MINUS_13       TimeZoneEnum = "Etc__GMT_minus_13"
	TIMEZONEENUM_PACIFIC__KIRITIMATI     TimeZoneEnum = "Pacific__Kiritimati"
	TIMEZONEENUM_ETC__GMT_MINUS_14       TimeZoneEnum = "Etc__GMT_minus_14"
)

// All allowed values of TimeZoneEnum enum
var AllowedTimeZoneEnumEnumValues = []TimeZoneEnum{
	"Etc__GMT_plus_12",
	"US__Samoa",
	"Etc__GMT_plus_11",
	"America__Atka",
	"US__Hawaii",
	"Etc__GMT_plus_10",
	"Pacific__Marquesas",
	"US__Alaska",
	"Pacific__Gambier",
	"Etc__GMT_plus_9",
	"PST8PDT",
	"Pacific__Pitcairn",
	"US__Pacific",
	"Etc__GMT_plus_8",
	"Mexico__BajaSur",
	"America__Boise",
	"America__Phoenix",
	"MST7MDT",
	"Etc__GMT_plus_7",
	"CST6CDT",
	"America__Chicago",
	"Canada__Saskatchewan",
	"America__Bahia_Banderas",
	"Etc__GMT_plus_6",
	"Chile__EasterIsland",
	"America__Bogota",
	"America__New_York",
	"EST5EDT",
	"America__Havana",
	"Etc__GMT_plus_5",
	"America__Caracas",
	"America__Cuiaba",
	"America__Santo_Domingo",
	"Canada__Atlantic",
	"America__Asuncion",
	"Etc__GMT_plus_4",
	"Canada__Newfoundland",
	"Chile__Continental",
	"Brazil__East",
	"America__Godthab",
	"America__Miquelon",
	"America__Buenos_Aires",
	"Etc__GMT_plus_3",
	"America__Noronha",
	"Etc__GMT_plus_2",
	"America__Scoresbysund",
	"Atlantic__Cape_Verde",
	"Etc__GMT_plus_1",
	"UTC",
	"Europe__London",
	"Africa__Casablanca",
	"Atlantic__Reykjavik",
	"Antarctica__Troll",
	"Europe__Paris",
	"Europe__Sarajevo",
	"Europe__Belgrade",
	"Europe__Rome",
	"Africa__Tunis",
	"Etc__GMT_minus_1",
	"Asia__Gaza",
	"Europe__Bucharest",
	"Europe__Helsinki",
	"Asia__Beirut",
	"Africa__Harare",
	"Asia__Damascus",
	"Asia__Amman",
	"Europe__Tiraspol",
	"Asia__Jerusalem",
	"Etc__GMT_minus_2",
	"Asia__Baghdad",
	"Africa__Asmera",
	"Etc__GMT_minus_3",
	"Asia__Tehran",
	"Asia__Baku",
	"Etc__GMT_minus_4",
	"Asia__Kabul",
	"Asia__Karachi",
	"Etc__GMT_minus_5",
	"Asia__Kolkata",
	"Asia__Katmandu",
	"Asia__Almaty",
	"Etc__GMT_minus_6",
	"Asia__Rangoon",
	"Asia__Hovd",
	"Asia__Bangkok",
	"Etc__GMT_minus_7",
	"Asia__Hong_Kong",
	"Asia__Brunei",
	"Asia__Singapore",
	"Etc__GMT_minus_8",
	"Asia__Pyongyang",
	"Australia__Eucla",
	"Asia__Seoul",
	"Etc__GMT_minus_9",
	"Australia__Darwin",
	"Australia__Adelaide",
	"Australia__Sydney",
	"Australia__Brisbane",
	"Asia__Magadan",
	"Etc__GMT_minus_10",
	"Australia__Lord_Howe",
	"Etc__GMT_minus_11",
	"Asia__Kamchatka",
	"Pacific__Fiji",
	"Antarctica__South_Pole",
	"Etc__GMT_minus_12",
	"Pacific__Chatham",
	"Pacific__Tongatapu",
	"Pacific__Apia",
	"Etc__GMT_minus_13",
	"Pacific__Kiritimati",
	"Etc__GMT_minus_14",
}

func (v *TimeZoneEnum) Value() string {
	return string(*v)
}
