package hll

import (
	"fmt"
	"slices"
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

type WeaponIdentifier string

const (
	WEAPON_M1903_SPRINGFIELD                            WeaponIdentifier = "M1903 SPRINGFIELD"
	WEAPON_FG42_X4                                      WeaponIdentifier = "FG42 x4"
	WEAPON_COAXIAL_M1919_SHERMAN_M4A3_75_W              WeaponIdentifier = "COAXIAL M1919 [Sherman M4A3(75)W]"
	WEAPON_OPEL_BLITZ_TRANSPORT                         WeaponIdentifier = "Opel Blitz (Transport)"
	WEAPON_PPSH_41                                      WeaponIdentifier = "PPSH 41"
	WEAPON_M1_CARBINE                                   WeaponIdentifier = "M1 CARBINE"
	WEAPON_M1918A2_BAR                                  WeaponIdentifier = "M1918A2 BAR"
	WEAPON_RIFLE_NO_4_MK_I_SNIPER                       WeaponIdentifier = "Rifle No.4 Mk I Sniper"
	WEAPON_M24_STIELHANDGRANATE                         WeaponIdentifier = "M24 STIELHANDGRANATE"
	WEAPON_PPSH_41_W_DRUM                               WeaponIdentifier = "PPSH 41 W/DRUM"
	WEAPON_HULL_M1919                                   WeaponIdentifier = "HULL M1919"
	WEAPON_76MM_ZIS_5_T34_76                            WeaponIdentifier = "76MM ZiS-5 [T34/76]"
	WEAPON_BOMBING_RUN                                  WeaponIdentifier = "BOMBING RUN"
	WEAPON_M1_GARAND                                    WeaponIdentifier = "M1 GARAND"
	WEAPON_FIRESPOT                                     WeaponIdentifier = "FireSpot"
	WEAPON_HULL_M1919_STUART_M5A1                       WeaponIdentifier = "HULL M1919 [Stuart M5A1]"
	WEAPON_SD_KFZ_234_PUMA                              WeaponIdentifier = "Sd.Kfz.234 Puma"
	WEAPON_GMC_CCKW_353_TRANSPORT                       WeaponIdentifier = "GMC CCKW 353 (Transport)"
	WEAPON_COAXIAL_M1919                                WeaponIdentifier = "COAXIAL M1919"
	WEAPON_LEWIS_GUN                                    WeaponIdentifier = "Lewis Gun"
	WEAPON_FLARE_GUN                                    WeaponIdentifier = "FLARE GUN"
	WEAPON_M97_TRENCH_GUN                               WeaponIdentifier = "M97 TRENCH GUN"
	WEAPON_M2_BROWNING_M3_HALF_TRACK                    WeaponIdentifier = "M2 Browning [M3 Half-track]"
	WEAPON_TELLERMINE_43                                WeaponIdentifier = "TELLERMINE 43"
	WEAPON_COAXIAL_MG34_SD_KFZ_234_PUMA                 WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.234 Puma]"
	WEAPON_BEDFORD_OYD_TRANSPORT                        WeaponIdentifier = "Bedford OYD (Transport)"
	WEAPON_COAXIAL_BESA_DAIMLER                         WeaponIdentifier = "COAXIAL BESA [Daimler]"
	WEAPON_75MM_M3_GUN                                  WeaponIdentifier = "75MM M3 GUN"
	WEAPON_COAXIAL_BESA                                 WeaponIdentifier = "COAXIAL BESA"
	WEAPON_PRECISION_STRIKE                             WeaponIdentifier = "PRECISION STRIKE"
	WEAPON_QF_25_POUNDER                                WeaponIdentifier = "QF 25-POUNDER"
	WEAPON_SCOPED_MOSIN_NAGANT_91_30                    WeaponIdentifier = "SCOPED MOSIN NAGANT 91/30"
	WEAPON_HULL_M1919_SHERMAN_M4A3_75_W                 WeaponIdentifier = "HULL M1919 [Sherman M4A3(75)W]"
	WEAPON_150MM_HOWITZER_SFH_18                        WeaponIdentifier = "150MM HOWITZER [sFH 18]"
	WEAPON_QF_2_POUNDER_DAIMLER                         WeaponIdentifier = "QF 2-POUNDER [Daimler]"
	WEAPON_HULL_M1919_SHERMAN_M4A3E2                    WeaponIdentifier = "HULL M1919 [Sherman M4A3E2]"
	WEAPON_POMZ_AP_MINE                                 WeaponIdentifier = "POMZ AP MINE"
	WEAPON_STEN_GUN_MK_II                               WeaponIdentifier = "Sten Gun Mk.II"
	WEAPON_M1928A1_THOMPSON                             WeaponIdentifier = "M1928A1 THOMPSON"
	WEAPON_CHURCHILL_MK_III                             WeaponIdentifier = "Churchill Mk.III"
	WEAPON_COAXIAL_MG34_SD_KFZ_171_PANTHER              WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.171 Panther]"
	WEAPON_COAXIAL_M1919_M3_STUART_HONEY                WeaponIdentifier = "COAXIAL M1919 [M3 Stuart Honey]"
	WEAPON_37MM_CANNON                                  WeaponIdentifier = "37MM CANNON"
	WEAPON_STUART_M5A1                                  WeaponIdentifier = "Stuart M5A1"
	WEAPON_M4A3_105MM                                   WeaponIdentifier = "M4A3 (105mm)"
	WEAPON_19_K_45MM                                    WeaponIdentifier = "19-K 45MM"
	WEAPON_HULL_DT_T34_76                               WeaponIdentifier = "HULL DT [T34/76]"
	WEAPON_HULL_DT_IS_1                                 WeaponIdentifier = "HULL DT [IS-1]"
	WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E    WeaponIdentifier = "HULL BESA 7.92mm [Churchill Mk III A.V.R.E.]"
	WEAPON_M1A1_THOMPSON                                WeaponIdentifier = "M1A1 THOMPSON"
	WEAPON_PANZER_III_AUSF_N                            WeaponIdentifier = "Panzer III Ausf.N"
	WEAPON_COAXIAL_MG34                                 WeaponIdentifier = "COAXIAL MG34"
	WEAPON_OQF_75MM                                     WeaponIdentifier = "OQF 75MM"
	WEAPON_19_K_45MM_BA_10                              WeaponIdentifier = "19-K 45MM [BA-10]"
	WEAPON_COAXIAL_BESA_CROMWELL                        WeaponIdentifier = "COAXIAL BESA [Cromwell]"
	WEAPON_76MM_ZIS_5                                   WeaponIdentifier = "76MM ZiS-5"
	WEAPON_NO_2_MK_5_FLARE_PISTOL                       WeaponIdentifier = "No.2 Mk 5 Flare Pistol"
	WEAPON_HULL_BESA                                    WeaponIdentifier = "HULL BESA"
	WEAPON_FLAMMENWERFER_41                             WeaponIdentifier = "FLAMMENWERFER 41"
	WEAPON_STG44                                        WeaponIdentifier = "STG44"
	WEAPON_GEWEHR_43                                    WeaponIdentifier = "GEWEHR 43"
	WEAPON_STRAFING_RUN                                 WeaponIdentifier = "STRAFING RUN"
	WEAPON_M43_STIELHANDGRANATE                         WeaponIdentifier = "M43 STIELHANDGRANATE"
	WEAPON_HULL_DT                                      WeaponIdentifier = "HULL DT"
	WEAPON_CHURCHILL_MK_III_A_V_R_E                     WeaponIdentifier = "Churchill Mk III A.V.R.E."
	WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III            WeaponIdentifier = "HULL BESA 7.92mm [Churchill Mk.III]"
	WEAPON_Satchel                                      WeaponIdentifier = "Satchel"
	WEAPON_M2_AP_MINE                                   WeaponIdentifier = "M2 AP MINE"
	WEAPON_COAXIAL_BESA_7_92MM                          WeaponIdentifier = "COAXIAL BESA 7.92mm"
	WEAPON_FAIRBAIRN_SYKES                              WeaponIdentifier = "Fairbairn–Sykes"
	WEAPON_HULL_BESA_7_92MM                             WeaponIdentifier = "HULL BESA 7.92mm"
	WEAPON_50MM_KWK_39_1_SD_KFZ_234_PUMA                WeaponIdentifier = "50mm KwK 39/1 [Sd.Kfz.234 Puma]"
	WEAPON_TETRARCH                                     WeaponIdentifier = "Tetrarch"
	WEAPON_T34_76                                       WeaponIdentifier = "T34/76"
	WEAPON_D_5T_85MM_IS_1                               WeaponIdentifier = "D-5T 85MM [IS-1]"
	WEAPON_STUH_43_L_12_STURMPANZER_IV                  WeaponIdentifier = "StuH 43 L/12 [Sturmpanzer IV]"
	WEAPON_122MM_HOWITZER                               WeaponIdentifier = "122MM HOWITZER"
	WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III         WeaponIdentifier = "COAXIAL BESA 7.92mm [Churchill Mk.III]"
	WEAPON_GAZ_67                                       WeaponIdentifier = "GAZ-67"
	WEAPON_A_P_SHRAPNEL_MINE_MK_II                      WeaponIdentifier = "A.P. Shrapnel Mine Mk II"
	WEAPON_MG34                                         WeaponIdentifier = "MG34"
	WEAPON_FLAMETHROWER                                 WeaponIdentifier = "FLAMETHROWER"
	WEAPON_SATCHEL_CHARGE                               WeaponIdentifier = "SATCHEL CHARGE"
	WEAPON_TOKAREV_TT33                                 WeaponIdentifier = "TOKAREV TT33"
	WEAPON_DAIMLER                                      WeaponIdentifier = "Daimler"
	WEAPON_FG42                                         WeaponIdentifier = "FG42"
	WEAPON_M2_FLAMETHROWER                              WeaponIdentifier = "M2 FLAMETHROWER"
	WEAPON_MOLOTOV                                      WeaponIdentifier = "MOLOTOV"
	WEAPON_45MM_M1937_T70                               WeaponIdentifier = "45MM M1937 [T70]"
	WEAPON_SHERMAN_M4A3E2                               WeaponIdentifier = "Sherman M4A3E2"
	WEAPON_QF_25_POUNDER_QF_25_POUNDER                  WeaponIdentifier = "QF 25-POUNDER [QF 25-Pounder]"
	WEAPON_BREN_GUN                                     WeaponIdentifier = "Bren Gun"
	WEAPON_MG_42                                        WeaponIdentifier = "MG 42"
	WEAPON_SCOPED_SVT40                                 WeaponIdentifier = "SCOPED SVT40"
	WEAPON_OQF_57MM                                     WeaponIdentifier = "OQF 57MM"
	WEAPON_SD_KFZ_171_PANTHER                           WeaponIdentifier = "Sd.Kfz.171 Panther"
	WEAPON_75MM_CANNON_PAK_40                           WeaponIdentifier = "75MM CANNON [PAK 40]"
	WEAPON_BA_10                                        WeaponIdentifier = "BA-10"
	WEAPON_OPEL_BLITZ_SUPPLY                            WeaponIdentifier = "Opel Blitz (Supply)"
	WEAPON_57MM_CANNON_M1_57MM                          WeaponIdentifier = "57MM CANNON [M1 57mm]"
	WEAPON_SD_KFZ_161_PANZER_IV                         WeaponIdentifier = "Sd.Kfz.161 Panzer IV"
	WEAPON_KUBELWAGEN                                   WeaponIdentifier = "Kubelwagen"
	WEAPON_MP40                                         WeaponIdentifier = "MP40"
	WEAPON_230MM_PETARD_CHURCHILL_MK_III_A_V_R_E        WeaponIdentifier = "230MM PETARD [Churchill Mk III A.V.R.E.]"
	WEAPON_QF_25_POUNDER_BISHOP_SP_25PDR                WeaponIdentifier = "QF 25 POUNDER [Bishop SP 25pdr]"
	WEAPON_COAXIAL_M1919_FIREFLY                        WeaponIdentifier = "COAXIAL M1919 [Firefly]"
	WEAPON_20MM_KWK_30                                  WeaponIdentifier = "20MM KWK 30"
	WEAPON_IS_1                                         WeaponIdentifier = "IS-1"
	WEAPON_OQF_57MM_CRUSADER_MK_III                     WeaponIdentifier = "OQF 57MM [Crusader Mk.III]"
	WEAPON_SD_KFZ_181_TIGER_1                           WeaponIdentifier = "Sd.Kfz.181 Tiger 1"
	WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2_76              WeaponIdentifier = "COAXIAL M1919 [Sherman M4A3E2(76)]"
	WEAPON_OQF_57MM_CHURCHILL_MK_III                    WeaponIdentifier = "OQF 57MM [Churchill Mk.III]"
	WEAPON_UNKNOWN                                      WeaponIdentifier = "UNKNOWN"
	WEAPON_MOSIN_NAGANT_91_30                           WeaponIdentifier = "MOSIN NAGANT 91/30"
	WEAPON_75MM_CANNON_SD_KFZ_161_PANZER_IV             WeaponIdentifier = "75MM CANNON [Sd.Kfz.161 Panzer IV]"
	WEAPON_COAXIAL_BESA_TETRARCH                        WeaponIdentifier = "COAXIAL BESA [Tetrarch]"
	WEAPON_57MM_CANNON                                  WeaponIdentifier = "57MM CANNON"
	WEAPON_OQF_75MM_CROMWELL                            WeaponIdentifier = "OQF 75MM [Cromwell]"
	WEAPON_M2_BROWNING                                  WeaponIdentifier = "M2 Browning"
	WEAPON_BROWNING_M1919                               WeaponIdentifier = "BROWNING M1919"
	WEAPON_20MM_KWK_30_SD_KFZ_121_LUCHS                 WeaponIdentifier = "20MM KWK 30 [Sd.Kfz.121 Luchs]"
	WEAPON_COAXIAL_M1919_M8_GREYHOUND                   WeaponIdentifier = "COAXIAL M1919 [M8 Greyhound]"
	WEAPON_PIAT                                         WeaponIdentifier = "PIAT"
	WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_VII            WeaponIdentifier = "HULL BESA 7.92mm [Churchill Mk.VII]"
	WEAPON_JEEP_WILLYS                                  WeaponIdentifier = "Jeep Willys"
	WEAPON_COAXIAL_DT                                   WeaponIdentifier = "COAXIAL DT"
	WEAPON_LUGER_P08                                    WeaponIdentifier = "LUGER P08"
	WEAPON_PANZERSCHRECK                                WeaponIdentifier = "PANZERSCHRECK"
	WEAPON_QF_25_POUNDER_GUN                            WeaponIdentifier = "QF 25 POUNDER"
	WEAPON_BISHOP_SP_25PDR                              WeaponIdentifier = "Bishop SP 25pdr"
	WEAPON_152MM_M_10T_KV_2                             WeaponIdentifier = "152MM M-10T [KV-2]"
	WEAPON_COAXIAL_MG34_SD_KFZ_161_PANZER_IV            WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.161 Panzer IV]"
	WEAPON_S_MINE                                       WeaponIdentifier = "S-MINE"
	WEAPON_HULL_MG34                                    WeaponIdentifier = "HULL MG34"
	WEAPON_HULL_MG34_SD_KFZ_161_PANZER_IV               WeaponIdentifier = "HULL MG34 [Sd.Kfz.161 Panzer IV]"
	WEAPON_ZIS_5_TRANSPORT                              WeaponIdentifier = "ZIS-5 (Transport)"
	WEAPON_155MM_HOWITZER                               WeaponIdentifier = "155MM HOWITZER"
	WEAPON_KARABINER_98K                                WeaponIdentifier = "KARABINER 98K"
	WEAPON_COAXIAL_DT_IS_1                              WeaponIdentifier = "COAXIAL DT [IS-1]"
	WEAPON_152MM_M_10T                                  WeaponIdentifier = "152MM M-10T"
	WEAPON_OQF_75MM_CHURCHILL_MK_VII                    WeaponIdentifier = "OQF 75MM [Churchill Mk.VII]"
	WEAPON_MOSIN_NAGANT_M38                             WeaponIdentifier = "MOSIN NAGANT M38"
	WEAPON_GMC_CCKW_353_SUPPLY                          WeaponIdentifier = "GMC CCKW 353 (Supply)"
	WEAPON_A_T_MINE_G_S_MK_V                            WeaponIdentifier = "A.T. Mine G.S. Mk V"
	WEAPON_88_KWK_36_L_56_SD_KFZ_181_TIGER_1            WeaponIdentifier = "88 KWK 36 L/56 [Sd.Kfz.181 Tiger 1]"
	WEAPON_75MM_CANNON_SHERMAN_M4A3_75_W                WeaponIdentifier = "75MM CANNON [Sherman M4A3(75)W]"
	WEAPON_QF_2_POUNDER                                 WeaponIdentifier = "QF 2-POUNDER"
	WEAPON_COAXIAL_BESA_CRUSADER_MK_III                 WeaponIdentifier = "COAXIAL BESA [Crusader Mk.III]"
	WEAPON_HULL_MG34_PANZER_III_AUSF_N                  WeaponIdentifier = "HULL MG34 [Panzer III Ausf.N]"
	WEAPON_KARABINER_98K_X8                             WeaponIdentifier = "KARABINER 98K x8"
	WEAPON_HULL_MG34_SD_KFZ_181_TIGER_1                 WeaponIdentifier = "HULL MG34 [Sd.Kfz.181 Tiger 1]"
	WEAPON_SHERMAN_M4A3_75_W                            WeaponIdentifier = "Sherman M4A3(75)W"
	WEAPON_WEBLEY_MK_VI                                 WeaponIdentifier = "Webley MK VI"
	WEAPON_NAGANT_M1895                                 WeaponIdentifier = "NAGANT M1895"
	WEAPON_QF_2_POUNDER_TETRARCH                        WeaponIdentifier = "QF 2-POUNDER [Tetrarch]"
	WEAPON_45MM_M1937                                   WeaponIdentifier = "45MM M1937"
	WEAPON_SMLE_NO_1_MK_III                             WeaponIdentifier = "SMLE No.1 Mk III"
	WEAPON_CROMWELL                                     WeaponIdentifier = "Cromwell"
	WEAPON_COAXIAL_M1919_STUART_M5A1                    WeaponIdentifier = "COAXIAL M1919 [Stuart M5A1]"
	WEAPON_T70                                          WeaponIdentifier = "T70"
	WEAPON_COAXIAL_DT_T70                               WeaponIdentifier = "COAXIAL DT [T70]"
	WEAPON_QF_17_POUNDER                                WeaponIdentifier = "QF 17-POUNDER"
	WEAPON_M3_KNIFE                                     WeaponIdentifier = "M3 KNIFE"
	WEAPON_88_KWK_36_L_56                               WeaponIdentifier = "88 KWK 36 L/56"
	WEAPON_SATCHEL                                      WeaponIdentifier = "SATCHEL"
	WEAPON_M6_37MM                                      WeaponIdentifier = "M6 37MM"
	WEAPON_HULL_MG34_SD_KFZ_171_PANTHER                 WeaponIdentifier = "HULL MG34 [Sd.Kfz.171 Panther]"
	WEAPON_DP_27                                        WeaponIdentifier = "DP-27"
	WEAPON_COAXIAL_MG34_SD_KFZ_181_TIGER_1              WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.181 Tiger 1]"
	WEAPON_SHERMAN_M4A3E2_76                            WeaponIdentifier = "Sherman M4A3E2(76)"
	WEAPON_MILLS_BOMB                                   WeaponIdentifier = "Mills Bomb"
	WEAPON_ZIS_5_SUPPLY                                 WeaponIdentifier = "ZIS-5 (Supply)"
	WEAPON_105MM_HOWITZER_M4A3_105MM                    WeaponIdentifier = "105MM HOWITZER [M4A3 (105mm)]"
	WEAPON_HULL_DT_KV_2                                 WeaponIdentifier = "HULL DT [KV-2]"
	WEAPON_37MM_CANNON_M3_STUART_HONEY                  WeaponIdentifier = "37MM CANNON [M3 Stuart Honey]"
	WEAPON_HULL_M1919_M3_STUART_HONEY                   WeaponIdentifier = "HULL M1919 [M3 Stuart Honey]"
	WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2                 WeaponIdentifier = "COAXIAL M1919 [Sherman M4A3E2]"
	WEAPON_HULL_M1919_SHERMAN_M4A3E2_76                 WeaponIdentifier = "HULL M1919 [Sherman M4A3E2(76)]"
	WEAPON_HULL_M1919_M4A3_105MM                        WeaponIdentifier = "HULL M1919 [M4A3 (105mm)]"
	WEAPON_155MM_HOWITZER_M114                          WeaponIdentifier = "155MM HOWITZER [M114]"
	WEAPON_COAXIAL_MG34_PANZER_III_AUSF_N               WeaponIdentifier = "COAXIAL MG34 [Panzer III Ausf.N]"
	WEAPON_7_5CM_KWK_37                                 WeaponIdentifier = "7.5CM KwK 37"
	WEAPON_QF_6_POUNDER                                 WeaponIdentifier = "QF 6-POUNDER"
	WEAPON_BAZOOKA                                      WeaponIdentifier = "BAZOOKA"
	WEAPON_SVT40                                        WeaponIdentifier = "SVT40"
	WEAPON_M8_GREYHOUND                                 WeaponIdentifier = "M8 Greyhound"
	WEAPON_122MM_HOWITZER_M1938_M_30                    WeaponIdentifier = "122MM HOWITZER [M1938 (M-30)]"
	WEAPON_BEDFORD_OYD_SUPPLY                           WeaponIdentifier = "Bedford OYD (Supply)"
	WEAPON_COAXIAL_MG34_SD_KFZ_121_LUCHS                WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.121 Luchs]"
	WEAPON_MPL_50_SPADE                                 WeaponIdentifier = "MPL-50 SPADE"
	WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E WeaponIdentifier = "COAXIAL BESA 7.92mm [Churchill Mk III A.V.R.E.]"
	WEAPON_M3_STUART_HONEY                              WeaponIdentifier = "M3 Stuart Honey"
	WEAPON_PTRS_41                                      WeaponIdentifier = "PTRS-41"
	WEAPON_57MM_CANNON_ZIS_2                            WeaponIdentifier = "57MM CANNON [ZiS-2]"
	WEAPON_STEN_GUN_MK_V                                WeaponIdentifier = "Sten Gun Mk.V"
	WEAPON_Unknown                                      WeaponIdentifier = "Unknown"
	WEAPON_75MM_M3_GUN_SHERMAN_M4A3E2                   WeaponIdentifier = "75MM M3 GUN [Sherman M4A3E2]"
	WEAPON_M3_GREASE_GUN                                WeaponIdentifier = "M3 GREASE GUN"
	WEAPON_QF_6_POUNDER_QF_6_POUNDER                    WeaponIdentifier = "QF 6-POUNDER [QF 6-Pounder]"
	WEAPON_NO_82_GRENADE                                WeaponIdentifier = "No.82 Grenade"
	WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_VII         WeaponIdentifier = "COAXIAL BESA 7.92mm [Churchill Mk.VII]"
	WEAPON_MG42                                         WeaponIdentifier = "MG42"
	WEAPON_D_5T_85MM                                    WeaponIdentifier = "D-5T 85MM"
	WEAPON_BOYS_ANTI_TANK_RIFLE                         WeaponIdentifier = "Boys Anti-tank Rifle"
	WEAPON_150MM_HOWITZER                               WeaponIdentifier = "150MM HOWITZER"
	WEAPON_RIFLE_NO_4_MK_I                              WeaponIdentifier = "Rifle No.4 Mk I"
	WEAPON_STEN_GUN                                     WeaponIdentifier = "Sten Gun"
	WEAPON_CHURCHILL_MK_VII                             WeaponIdentifier = "Churchill Mk.VII"
	WEAPON_TM_35_AT_MINE                                WeaponIdentifier = "TM-35 AT MINE"
	WEAPON_FELDSPATEN                                   WeaponIdentifier = "FELDSPATEN"
	WEAPON_50MM_KWK_39_1                                WeaponIdentifier = "50MM KWK 39/1"
	WEAPON_SD_KFZ_121_LUCHS                             WeaponIdentifier = "Sd.Kfz.121 Luchs"
	WEAPON_76MM_M1_GUN_SHERMAN_M4A3E2_76                WeaponIdentifier = "76MM M1 GUN [Sherman M4A3E2(76)]"
	WEAPON_MK2_GRENADE                                  WeaponIdentifier = "MK2 GRENADE"
	WEAPON_MG_42_SD_KFZ_251_HALF_TRACK                  WeaponIdentifier = "MG 42 [Sd.Kfz 251 Half-track]"
	WEAPON_76MM_M1_GUN                                  WeaponIdentifier = "76MM M1 GUN"
	WEAPON_M3_HALF_TRACK                                WeaponIdentifier = "M3 Half-track"
	WEAPON_COAXIAL_M1919_M4A3_105MM                     WeaponIdentifier = "COAXIAL M1919 [M4A3 (105mm)]"
	WEAPON_7_5CM_KWK_37_PANZER_III_AUSF_N               WeaponIdentifier = "7.5CM KwK 37 [Panzer III Ausf.N]"
	WEAPON_STURMPANZER_IV                               WeaponIdentifier = "Sturmpanzer IV"
	WEAPON_75MM_CANNON_SD_KFZ_171_PANTHER               WeaponIdentifier = "75MM CANNON [Sd.Kfz.171 Panther]"
	WEAPON_COAXIAL_DT_BA_10                             WeaponIdentifier = "COAXIAL DT [BA-10]"
	WEAPON_RG_42_GRENADE                                WeaponIdentifier = "RG-42 GRENADE"
	WEAPON_WALTHER_P38                                  WeaponIdentifier = "WALTHER P38"
	WEAPON_CRUSADER_MK_III                              WeaponIdentifier = "Crusader Mk.III"
	WEAPON_COLT_M1911                                   WeaponIdentifier = "COLT M1911"
	WEAPON_FIREFLY                                      WeaponIdentifier = "Firefly"
	WEAPON_75MM_CANNON                                  WeaponIdentifier = "75MM CANNON"
	WEAPON_KV_2                                         WeaponIdentifier = "KV-2"
	WEAPON_37MM_CANNON_STUART_M5A1                      WeaponIdentifier = "37MM CANNON [Stuart M5A1]"
	WEAPON_HULL_BESA_CROMWELL                           WeaponIdentifier = "HULL BESA [Cromwell]"
	WEAPON_QF_17_POUNDER_FIREFLY                        WeaponIdentifier = "QF 17-POUNDER [Firefly]"
	WEAPON_COAXIAL_DT_T34_76                            WeaponIdentifier = "COAXIAL DT [T34/76]"
	WEAPON_MOSIN_NAGANT_1891                            WeaponIdentifier = "MOSIN NAGANT 1891"
	WEAPON_M6_37MM_M8_GREYHOUND                         WeaponIdentifier = "M6 37mm [M8 Greyhound]"
	WEAPON_M1A1_AT_MINE                                 WeaponIdentifier = "M1A1 AT MINE"
	WEAPON_SD_KFZ_251_HALF_TRACK                        WeaponIdentifier = "Sd.Kfz 251 Half-track"
)

type WeaponType string

const (
	WEAPON_TYPE_ROCKET_LAUNCHER     WeaponType = "Rocket Launcher"
	WEAPON_TYPE_ASSAULT_RIFLE       WeaponType = "Assault Rifle"
	WEAPON_TYPE_ANTI_PERSONNEL_MINE WeaponType = "Anti-Personnel Mine"
	WEAPON_TYPE_SATCHEL             WeaponType = "Satchel"
	WEAPON_TYPE_UNKNOWN             WeaponType = "Unknown"
	WEAPON_TYPE_ANTI_TANK_MINE      WeaponType = "Anti-Tank Mine"
	WEAPON_TYPE_MOUNTED_MG          WeaponType = "Mounted MG"
	WEAPON_TYPE_ANTI_MATERIEL_RIFLE WeaponType = "Anti-Materiel Rifle"
	WEAPON_TYPE_REVOLVER            WeaponType = "Revolver"
	WEAPON_TYPE_SEMI_AUTO_RIFLE     WeaponType = "Semi-Auto Rifle"
	WEAPON_TYPE_TANK_CANNON         WeaponType = "Tank Cannon"
	WEAPON_TYPE_MELEE               WeaponType = "Melee"
	WEAPON_TYPE_ROADKILL            WeaponType = "Roadkill"
	WEAPON_TYPE_TANK_COAXIAL_MG     WeaponType = "Tank Coaxial MG"
	WEAPON_TYPE_FLAMETHROWER        WeaponType = "Flamethrower"
	WEAPON_TYPE_BOLT_ACTION_RIFLE   WeaponType = "Bolt Action Rifle"
	WEAPON_TYPE_COMMANDER_ABILITY   WeaponType = "Commander Ability"
	WEAPON_TYPE_ARTILLERY           WeaponType = "Artillery"
	WEAPON_TYPE_TANK_HULL_MG        WeaponType = "Tank Hull MG"
	WEAPON_TYPE_GRENADE             WeaponType = "Grenade"
	WEAPON_TYPE_PISTOL              WeaponType = "Pistol"
	WEAPON_TYPE_FLARE_GUN           WeaponType = "Flare Gun"
	WEAPON_TYPE_MACHINE_GUN         WeaponType = "Machine Gun"
	WEAPON_TYPE_SUBMACHINE_GUN      WeaponType = "Submachine Gun"
	WEAPON_TYPE_ANTI_TANK_GUN       WeaponType = "Anti-Tank Gun"
	WEAPON_TYPE_SHOTGUN             WeaponType = "Shotgun"
)

type Weapon struct {
	ID            WeaponIdentifier
	Name          string
	Type          WeaponType
	Factions      []FactionIdentifier
	Magnification int
}

var weaponMap = map[WeaponIdentifier]Weapon{
	WEAPON_M1_GARAND: {
		ID:            WEAPON_M1_GARAND,
		Name:          "M1 Garand",
		Type:          WEAPON_TYPE_SEMI_AUTO_RIFLE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M1_CARBINE: {
		ID:            WEAPON_M1_CARBINE,
		Name:          "M1 Carbine",
		Type:          WEAPON_TYPE_SEMI_AUTO_RIFLE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M1A1_THOMPSON: {
		ID:            WEAPON_M1A1_THOMPSON,
		Name:          "M1A1 Thompson",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M3_GREASE_GUN: {
		ID:            WEAPON_M3_GREASE_GUN,
		Name:          "M3 Grease Gun",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M1918A2_BAR: {
		ID:            WEAPON_M1918A2_BAR,
		Name:          "M1918A2 BAR",
		Type:          WEAPON_TYPE_ASSAULT_RIFLE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_BROWNING_M1919: {
		ID:            WEAPON_BROWNING_M1919,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_MACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M1903_SPRINGFIELD: {
		ID:            WEAPON_M1903_SPRINGFIELD,
		Name:          "M1903 Springfield",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 4,
	},
	WEAPON_M97_TRENCH_GUN: {
		ID:            WEAPON_M97_TRENCH_GUN,
		Name:          "M97 Trench Gun",
		Type:          WEAPON_TYPE_SHOTGUN,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COLT_M1911: {
		ID:            WEAPON_COLT_M1911,
		Name:          "Colt M1911",
		Type:          WEAPON_TYPE_PISTOL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M3_KNIFE: {
		ID:            WEAPON_M3_KNIFE,
		Name:          "M3 Knife",
		Type:          WEAPON_TYPE_MELEE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_SATCHEL: {
		ID:            WEAPON_SATCHEL,
		Name:          "Satchel Charge",
		Type:          WEAPON_TYPE_SATCHEL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_MK2_GRENADE: {
		ID:            WEAPON_MK2_GRENADE,
		Name:          "Mk 2 Grenade",
		Type:          WEAPON_TYPE_GRENADE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M2_FLAMETHROWER: {
		ID:            WEAPON_M2_FLAMETHROWER,
		Name:          "M2 Flamethrower",
		Type:          WEAPON_TYPE_FLAMETHROWER,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_BAZOOKA: {
		ID:            WEAPON_BAZOOKA,
		Name:          "Bazooka",
		Type:          WEAPON_TYPE_ROCKET_LAUNCHER,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M2_AP_MINE: {
		ID:            WEAPON_M2_AP_MINE,
		Name:          "M2 AP Mine",
		Type:          WEAPON_TYPE_ANTI_PERSONNEL_MINE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M1A1_AT_MINE: {
		ID:            WEAPON_M1A1_AT_MINE,
		Name:          "M1A1 AT Mine",
		Type:          WEAPON_TYPE_ANTI_TANK_MINE,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_FLARE_GUN: {
		ID:            WEAPON_FLARE_GUN,
		Name:          "Flare Gun",
		Type:          WEAPON_TYPE_FLARE_GUN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_SOV, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_57MM_CANNON_M1_57MM: {
		ID:            WEAPON_57MM_CANNON_M1_57MM,
		Name:          "57mm Cannon",
		Type:          WEAPON_TYPE_ANTI_TANK_GUN,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_155MM_HOWITZER_M114: {
		ID:            WEAPON_155MM_HOWITZER_M114,
		Name:          "155mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M8_GREYHOUND: {
		ID:            WEAPON_M8_GREYHOUND,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_STUART_M5A1: {
		ID:            WEAPON_STUART_M5A1,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_SHERMAN_M4A3_75_W: {
		ID:            WEAPON_SHERMAN_M4A3_75_W,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_SHERMAN_M4A3E2: {
		ID:            WEAPON_SHERMAN_M4A3E2,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_SHERMAN_M4A3E2_76: {
		ID:            WEAPON_SHERMAN_M4A3E2_76,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_GMC_CCKW_353_SUPPLY: {
		ID:            WEAPON_GMC_CCKW_353_SUPPLY,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_GMC_CCKW_353_TRANSPORT: {
		ID:            WEAPON_GMC_CCKW_353_TRANSPORT,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M3_HALF_TRACK: {
		ID:            WEAPON_M3_HALF_TRACK,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_SOV, FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_JEEP_WILLYS: {
		ID:            WEAPON_JEEP_WILLYS,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M4A3_105MM: {
		ID:            WEAPON_M4A3_105MM,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_CW},
		Magnification: 0,
	},
	WEAPON_M6_37MM_M8_GREYHOUND: {
		ID:            WEAPON_M6_37MM_M8_GREYHOUND,
		Name:          "37mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_M8_GREYHOUND: {
		ID:            WEAPON_COAXIAL_M1919_M8_GREYHOUND,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_37MM_CANNON_STUART_M5A1: {
		ID:            WEAPON_37MM_CANNON_STUART_M5A1,
		Name:          "37mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_STUART_M5A1: {
		ID:            WEAPON_COAXIAL_M1919_STUART_M5A1,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_HULL_M1919_STUART_M5A1: {
		ID:            WEAPON_HULL_M1919_STUART_M5A1,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_75MM_CANNON_SHERMAN_M4A3_75_W: {
		ID:            WEAPON_75MM_CANNON_SHERMAN_M4A3_75_W,
		Name:          "75mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_SHERMAN_M4A3_75_W: {
		ID:            WEAPON_COAXIAL_M1919_SHERMAN_M4A3_75_W,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_HULL_M1919_SHERMAN_M4A3_75_W: {
		ID:            WEAPON_HULL_M1919_SHERMAN_M4A3_75_W,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_75MM_M3_GUN_SHERMAN_M4A3E2: {
		ID:            WEAPON_75MM_M3_GUN_SHERMAN_M4A3E2,
		Name:          "75mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2: {
		ID:            WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_HULL_M1919_SHERMAN_M4A3E2: {
		ID:            WEAPON_HULL_M1919_SHERMAN_M4A3E2,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_76MM_M1_GUN_SHERMAN_M4A3E2_76: {
		ID:            WEAPON_76MM_M1_GUN_SHERMAN_M4A3E2_76,
		Name:          "76mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2_76: {
		ID:            WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2_76,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_HULL_M1919_SHERMAN_M4A3E2_76: {
		ID:            WEAPON_HULL_M1919_SHERMAN_M4A3E2_76,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_M2_BROWNING_M3_HALF_TRACK: {
		ID:            WEAPON_M2_BROWNING_M3_HALF_TRACK,
		Name:          "M2 Browning",
		Type:          WEAPON_TYPE_MOUNTED_MG,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_SOV, FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_105MM_HOWITZER_M4A3_105MM: {
		ID:            WEAPON_105MM_HOWITZER_M4A3_105MM,
		Name:          "105mm Howitzer",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_M4A3_105MM: {
		ID:            WEAPON_COAXIAL_M1919_M4A3_105MM,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_HULL_M1919_M4A3_105MM: {
		ID:            WEAPON_HULL_M1919_M4A3_105MM,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_57MM_CANNON: {
		ID:            WEAPON_57MM_CANNON,
		Name:          "57mm Cannon",
		Type:          WEAPON_TYPE_ANTI_TANK_GUN,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_155MM_HOWITZER: {
		ID:            WEAPON_155MM_HOWITZER,
		Name:          "155mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_76MM_M1_GUN: {
		ID:            WEAPON_76MM_M1_GUN,
		Name:          "76mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_75MM_M3_GUN: {
		ID:            WEAPON_75MM_M3_GUN,
		Name:          "75mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_75MM_CANNON: {
		ID:            WEAPON_75MM_CANNON,
		Name:          "75mm Cannon",
		Type:          WEAPON_TYPE_UNKNOWN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_37MM_CANNON: {
		ID:            WEAPON_37MM_CANNON,
		Name:          "37mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_CW},
		Magnification: 0,
	},
	WEAPON_M6_37MM: {
		ID:            WEAPON_M6_37MM,
		Name:          "M6 37mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_US},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919: {
		ID:            WEAPON_COAXIAL_M1919,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_CW},
		Magnification: 0,
	},
	WEAPON_HULL_M1919: {
		ID:            WEAPON_HULL_M1919,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_CW},
		Magnification: 0,
	},
	WEAPON_M2_BROWNING: {
		ID:            WEAPON_M2_BROWNING,
		Name:          "M2 Browning",
		Type:          WEAPON_TYPE_MOUNTED_MG,
		Factions:      []FactionIdentifier{FACTION_US, FACTION_SOV, FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_KARABINER_98K: {
		ID:            WEAPON_KARABINER_98K,
		Name:          "Karabiner 98k",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_GEWEHR_43: {
		ID:            WEAPON_GEWEHR_43,
		Name:          "G43",
		Type:          WEAPON_TYPE_SEMI_AUTO_RIFLE,
		Factions:      []FactionIdentifier{FACTION_GER},
		Magnification: 0,
	},
	WEAPON_STG44: {
		ID:            WEAPON_STG44,
		Name:          "STG44",
		Type:          WEAPON_TYPE_ASSAULT_RIFLE,
		Factions:      []FactionIdentifier{FACTION_GER},
		Magnification: 0,
	},
	WEAPON_FG42: {
		ID:            WEAPON_FG42,
		Name:          "FG42",
		Type:          WEAPON_TYPE_ASSAULT_RIFLE,
		Factions:      []FactionIdentifier{FACTION_GER},
		Magnification: 0,
	},
	WEAPON_MP40: {
		ID:            WEAPON_MP40,
		Name:          "MP40",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_MG34: {
		ID:            WEAPON_MG34,
		Name:          "MG34",
		Type:          WEAPON_TYPE_MACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_MG42: {
		ID:            WEAPON_MG42,
		Name:          "MG42",
		Type:          WEAPON_TYPE_MACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_FLAMMENWERFER_41: {
		ID:            WEAPON_FLAMMENWERFER_41,
		Name:          "Flammenwerfer 41",
		Type:          WEAPON_TYPE_FLAMETHROWER,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_KARABINER_98K_X8: {
		ID:            WEAPON_KARABINER_98K_X8,
		Name:          "Karabiner 98k",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 8,
	},
	WEAPON_FG42_X4: {
		ID:            WEAPON_FG42_X4,
		Name:          "FG42",
		Type:          WEAPON_TYPE_SEMI_AUTO_RIFLE,
		Factions:      []FactionIdentifier{FACTION_GER},
		Magnification: 4,
	},
	WEAPON_LUGER_P08: {
		ID:            WEAPON_LUGER_P08,
		Name:          "Luger P08",
		Type:          WEAPON_TYPE_PISTOL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_WALTHER_P38: {
		ID:            WEAPON_WALTHER_P38,
		Name:          "Walther P38",
		Type:          WEAPON_TYPE_PISTOL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_FELDSPATEN: {
		ID:            WEAPON_FELDSPATEN,
		Name:          "Feldspaten",
		Type:          WEAPON_TYPE_MELEE,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_M24_STIELHANDGRANATE: {
		ID:            WEAPON_M24_STIELHANDGRANATE,
		Name:          "M24 Stielhandgranate",
		Type:          WEAPON_TYPE_GRENADE,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_M43_STIELHANDGRANATE: {
		ID:            WEAPON_M43_STIELHANDGRANATE,
		Name:          "M43 Stielhandgranate",
		Type:          WEAPON_TYPE_GRENADE,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_PANZERSCHRECK: {
		ID:            WEAPON_PANZERSCHRECK,
		Name:          "Panzerschreck",
		Type:          WEAPON_TYPE_ROCKET_LAUNCHER,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_S_MINE: {
		ID:            WEAPON_S_MINE,
		Name:          "S-Mine",
		Type:          WEAPON_TYPE_ANTI_PERSONNEL_MINE,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_TELLERMINE_43: {
		ID:            WEAPON_TELLERMINE_43,
		Name:          "Tellermine 43",
		Type:          WEAPON_TYPE_ANTI_TANK_MINE,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_75MM_CANNON_PAK_40: {
		ID:            WEAPON_75MM_CANNON_PAK_40,
		Name:          "75mm Cannon",
		Type:          WEAPON_TYPE_ANTI_TANK_GUN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_150MM_HOWITZER_SFH_18: {
		ID:            WEAPON_150MM_HOWITZER_SFH_18,
		Name:          "150mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_SD_KFZ_234_PUMA: {
		ID:            WEAPON_SD_KFZ_234_PUMA,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_SD_KFZ_121_LUCHS: {
		ID:            WEAPON_SD_KFZ_121_LUCHS,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_SD_KFZ_161_PANZER_IV: {
		ID:            WEAPON_SD_KFZ_161_PANZER_IV,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_SD_KFZ_171_PANTHER: {
		ID:            WEAPON_SD_KFZ_171_PANTHER,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_SD_KFZ_181_TIGER_1: {
		ID:            WEAPON_SD_KFZ_181_TIGER_1,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_OPEL_BLITZ_SUPPLY: {
		ID:            WEAPON_OPEL_BLITZ_SUPPLY,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_OPEL_BLITZ_TRANSPORT: {
		ID:            WEAPON_OPEL_BLITZ_TRANSPORT,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_SD_KFZ_251_HALF_TRACK: {
		ID:            WEAPON_SD_KFZ_251_HALF_TRACK,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_KUBELWAGEN: {
		ID:            WEAPON_KUBELWAGEN,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_STURMPANZER_IV: {
		ID:            WEAPON_STURMPANZER_IV,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER},
		Magnification: 0,
	},
	WEAPON_PANZER_III_AUSF_N: {
		ID:            WEAPON_PANZER_III_AUSF_N,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_GER},
		Magnification: 0,
	},
	WEAPON_50MM_KWK_39_1_SD_KFZ_234_PUMA: {
		ID:            WEAPON_50MM_KWK_39_1_SD_KFZ_234_PUMA,
		Name:          "50mm KwK 39/1",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_COAXIAL_MG34_SD_KFZ_234_PUMA: {
		ID:            WEAPON_COAXIAL_MG34_SD_KFZ_234_PUMA,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_20MM_KWK_30_SD_KFZ_121_LUCHS: {
		ID:            WEAPON_20MM_KWK_30_SD_KFZ_121_LUCHS,
		Name:          "20mm KwK 30",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_COAXIAL_MG34_SD_KFZ_121_LUCHS: {
		ID:            WEAPON_COAXIAL_MG34_SD_KFZ_121_LUCHS,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_75MM_CANNON_SD_KFZ_161_PANZER_IV: {
		ID:            WEAPON_75MM_CANNON_SD_KFZ_161_PANZER_IV,
		Name:          "75mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_COAXIAL_MG34_SD_KFZ_161_PANZER_IV: {
		ID:            WEAPON_COAXIAL_MG34_SD_KFZ_161_PANZER_IV,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_HULL_MG34_SD_KFZ_161_PANZER_IV: {
		ID:            WEAPON_HULL_MG34_SD_KFZ_161_PANZER_IV,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_75MM_CANNON_SD_KFZ_171_PANTHER: {
		ID:            WEAPON_75MM_CANNON_SD_KFZ_171_PANTHER,
		Name:          "75mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_COAXIAL_MG34_SD_KFZ_171_PANTHER: {
		ID:            WEAPON_COAXIAL_MG34_SD_KFZ_171_PANTHER,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_HULL_MG34_SD_KFZ_171_PANTHER: {
		ID:            WEAPON_HULL_MG34_SD_KFZ_171_PANTHER,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_88_KWK_36_L_56_SD_KFZ_181_TIGER_1: {
		ID:            WEAPON_88_KWK_36_L_56_SD_KFZ_181_TIGER_1,
		Name:          "88mm KwK 36 L/56",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_COAXIAL_MG34_SD_KFZ_181_TIGER_1: {
		ID:            WEAPON_COAXIAL_MG34_SD_KFZ_181_TIGER_1,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_HULL_MG34_SD_KFZ_181_TIGER_1: {
		ID:            WEAPON_HULL_MG34_SD_KFZ_181_TIGER_1,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_MG_42_SD_KFZ_251_HALF_TRACK: {
		ID:            WEAPON_MG_42_SD_KFZ_251_HALF_TRACK,
		Name:          "MG42",
		Type:          WEAPON_TYPE_MOUNTED_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_STUH_43_L_12_STURMPANZER_IV: {
		ID:            WEAPON_STUH_43_L_12_STURMPANZER_IV,
		Name:          "StuH 43 L/12",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER},
		Magnification: 0,
	},
	WEAPON_7_5CM_KWK_37_PANZER_III_AUSF_N: {
		ID:            WEAPON_7_5CM_KWK_37_PANZER_III_AUSF_N,
		Name:          "75mm KwK 37",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_COAXIAL_MG34_PANZER_III_AUSF_N: {
		ID:            WEAPON_COAXIAL_MG34_PANZER_III_AUSF_N,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_HULL_MG34_PANZER_III_AUSF_N: {
		ID:            WEAPON_HULL_MG34_PANZER_III_AUSF_N,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_150MM_HOWITZER: {
		ID:            WEAPON_150MM_HOWITZER,
		Name:          "150mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_50MM_KWK_39_1: {
		ID:            WEAPON_50MM_KWK_39_1,
		Name:          "50mm KwK 39/1",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_20MM_KWK_30: {
		ID:            WEAPON_20MM_KWK_30,
		Name:          "20mm KwK 30",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_88_KWK_36_L_56: {
		ID:            WEAPON_88_KWK_36_L_56,
		Name:          "88mm KwK 36 L/56",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_COAXIAL_MG34: {
		ID:            WEAPON_COAXIAL_MG34,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_HULL_MG34: {
		ID:            WEAPON_HULL_MG34,
		Name:          "MG34",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_MG_42: {
		ID:            WEAPON_MG_42,
		Name:          "MG42",
		Type:          WEAPON_TYPE_MOUNTED_MG,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_7_5CM_KWK_37: {
		ID:            WEAPON_7_5CM_KWK_37,
		Name:          "75mm KwK 37",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_DAK},
		Magnification: 0,
	},
	WEAPON_MOSIN_NAGANT_1891: {
		ID:            WEAPON_MOSIN_NAGANT_1891,
		Name:          "Mosin-Nagant 1891",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_MOSIN_NAGANT_91_30: {
		ID:            WEAPON_MOSIN_NAGANT_91_30,
		Name:          "Mosin-Nagant 91/30",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_MOSIN_NAGANT_M38: {
		ID:            WEAPON_MOSIN_NAGANT_M38,
		Name:          "Mosin-Nagant M38",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_SVT40: {
		ID:            WEAPON_SVT40,
		Name:          "SVT-40",
		Type:          WEAPON_TYPE_SEMI_AUTO_RIFLE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_PPSH_41: {
		ID:            WEAPON_PPSH_41,
		Name:          "PPSh-41",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_PPSH_41_W_DRUM: {
		ID:            WEAPON_PPSH_41_W_DRUM,
		Name:          "PPSh-41 with Drum",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_DP_27: {
		ID:            WEAPON_DP_27,
		Name:          "DP-27",
		Type:          WEAPON_TYPE_MACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_SCOPED_MOSIN_NAGANT_91_30: {
		ID:            WEAPON_SCOPED_MOSIN_NAGANT_91_30,
		Name:          "Mosin-Nagant 91/30",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 4,
	},
	WEAPON_SCOPED_SVT40: {
		ID:            WEAPON_SCOPED_SVT40,
		Name:          "SVT-40",
		Type:          WEAPON_TYPE_SEMI_AUTO_RIFLE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 4,
	},
	WEAPON_NAGANT_M1895: {
		ID:            WEAPON_NAGANT_M1895,
		Name:          "Nagant M1895",
		Type:          WEAPON_TYPE_REVOLVER,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_TOKAREV_TT33: {
		ID:            WEAPON_TOKAREV_TT33,
		Name:          "Tokarev TT-33",
		Type:          WEAPON_TYPE_PISTOL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_MPL_50_SPADE: {
		ID:            WEAPON_MPL_50_SPADE,
		Name:          "MPL-50 Spade",
		Type:          WEAPON_TYPE_MELEE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_SATCHEL_CHARGE: {
		ID:            WEAPON_SATCHEL_CHARGE,
		Name:          "Satchel Charge",
		Type:          WEAPON_TYPE_SATCHEL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_RG_42_GRENADE: {
		ID:            WEAPON_RG_42_GRENADE,
		Name:          "RG-42 Grenade",
		Type:          WEAPON_TYPE_GRENADE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_MOLOTOV: {
		ID:            WEAPON_MOLOTOV,
		Name:          "Molotov",
		Type:          WEAPON_TYPE_GRENADE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_PTRS_41: {
		ID:            WEAPON_PTRS_41,
		Name:          "PTRS-41",
		Type:          WEAPON_TYPE_ANTI_MATERIEL_RIFLE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_POMZ_AP_MINE: {
		ID:            WEAPON_POMZ_AP_MINE,
		Name:          "POMZ AP Mine",
		Type:          WEAPON_TYPE_ANTI_PERSONNEL_MINE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_TM_35_AT_MINE: {
		ID:            WEAPON_TM_35_AT_MINE,
		Name:          "TM-35 AT Mine",
		Type:          WEAPON_TYPE_ANTI_TANK_MINE,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_57MM_CANNON_ZIS_2: {
		ID:            WEAPON_57MM_CANNON_ZIS_2,
		Name:          "57mm Cannon",
		Type:          WEAPON_TYPE_ANTI_TANK_GUN,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_122MM_HOWITZER_M1938_M_30: {
		ID:            WEAPON_122MM_HOWITZER_M1938_M_30,
		Name:          "122mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_BA_10: {
		ID:            WEAPON_BA_10,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_T70: {
		ID:            WEAPON_T70,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_T34_76: {
		ID:            WEAPON_T34_76,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_IS_1: {
		ID:            WEAPON_IS_1,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_ZIS_5_SUPPLY: {
		ID:            WEAPON_ZIS_5_SUPPLY,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_ZIS_5_TRANSPORT: {
		ID:            WEAPON_ZIS_5_TRANSPORT,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_KV_2: {
		ID:            WEAPON_KV_2,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_GAZ_67: {
		ID:            WEAPON_GAZ_67,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_19_K_45MM_BA_10: {
		ID:            WEAPON_19_K_45MM_BA_10,
		Name:          "45mm M1932",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_COAXIAL_DT_BA_10: {
		ID:            WEAPON_COAXIAL_DT_BA_10,
		Name:          "DT",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_45MM_M1937_T70: {
		ID:            WEAPON_45MM_M1937_T70,
		Name:          "45mm M1937",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_COAXIAL_DT_T70: {
		ID:            WEAPON_COAXIAL_DT_T70,
		Name:          "DT",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_76MM_ZIS_5_T34_76: {
		ID:            WEAPON_76MM_ZIS_5_T34_76,
		Name:          "76mm M1940",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_COAXIAL_DT_T34_76: {
		ID:            WEAPON_COAXIAL_DT_T34_76,
		Name:          "DT",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_HULL_DT_T34_76: {
		ID:            WEAPON_HULL_DT_T34_76,
		Name:          "DT",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_D_5T_85MM_IS_1: {
		ID:            WEAPON_D_5T_85MM_IS_1,
		Name:          "D-5T 85mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_COAXIAL_DT_IS_1: {
		ID:            WEAPON_COAXIAL_DT_IS_1,
		Name:          "DT",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_HULL_DT_IS_1: {
		ID:            WEAPON_HULL_DT_IS_1,
		Name:          "DT",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_152MM_M_10T_KV_2: {
		ID:            WEAPON_152MM_M_10T_KV_2,
		Name:          "M-10T 152mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_HULL_DT_KV_2: {
		ID:            WEAPON_HULL_DT_KV_2,
		Name:          "DT",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_122MM_HOWITZER: {
		ID:            WEAPON_122MM_HOWITZER,
		Name:          "122mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_19_K_45MM: {
		ID:            WEAPON_19_K_45MM,
		Name:          "45mm M1932",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_45MM_M1937: {
		ID:            WEAPON_45MM_M1937,
		Name:          "45mm M1937",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_76MM_ZIS_5: {
		ID:            WEAPON_76MM_ZIS_5,
		Name:          "76mm M1940",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_D_5T_85MM: {
		ID:            WEAPON_D_5T_85MM,
		Name:          "D-5T 85mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_COAXIAL_DT: {
		ID:            WEAPON_COAXIAL_DT,
		Name:          "COAXIAL DT",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_HULL_DT: {
		ID:            WEAPON_HULL_DT,
		Name:          "HULL DT",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_152MM_M_10T: {
		ID:            WEAPON_152MM_M_10T,
		Name:          "M-10T 152mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_SOV},
		Magnification: 0,
	},
	WEAPON_SMLE_NO_1_MK_III: {
		ID:            WEAPON_SMLE_NO_1_MK_III,
		Name:          "SMLE Mk III",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_RIFLE_NO_4_MK_I: {
		ID:            WEAPON_RIFLE_NO_4_MK_I,
		Name:          "No.4 Rifle Mk I",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_STEN_GUN: {
		ID:            WEAPON_STEN_GUN,
		Name:          "Sten Mk II",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_STEN_GUN_MK_II: {
		ID:            WEAPON_STEN_GUN_MK_II,
		Name:          "Sten Mk II",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_STEN_GUN_MK_V: {
		ID:            WEAPON_STEN_GUN_MK_V,
		Name:          "Sten Mk V",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_M1928A1_THOMPSON: {
		ID:            WEAPON_M1928A1_THOMPSON,
		Name:          "M1928A1 Thompson",
		Type:          WEAPON_TYPE_SUBMACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_BREN_GUN: {
		ID:            WEAPON_BREN_GUN,
		Name:          "Bren Gun",
		Type:          WEAPON_TYPE_ASSAULT_RIFLE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_LEWIS_GUN: {
		ID:            WEAPON_LEWIS_GUN,
		Name:          "Lewis Gun",
		Type:          WEAPON_TYPE_MACHINE_GUN,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_FLAMETHROWER: {
		ID:            WEAPON_FLAMETHROWER,
		Name:          "Lifebuoy Flamethrower",
		Type:          WEAPON_TYPE_FLAMETHROWER,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_RIFLE_NO_4_MK_I_SNIPER: {
		ID:            WEAPON_RIFLE_NO_4_MK_I_SNIPER,
		Name:          "No.4 Rifle Mk I",
		Type:          WEAPON_TYPE_BOLT_ACTION_RIFLE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 8,
	},
	WEAPON_WEBLEY_MK_VI: {
		ID:            WEAPON_WEBLEY_MK_VI,
		Name:          "Webley Mk IV",
		Type:          WEAPON_TYPE_REVOLVER,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_FAIRBAIRN_SYKES: {
		ID:            WEAPON_FAIRBAIRN_SYKES,
		Name:          "Fairbairn-Sykes",
		Type:          WEAPON_TYPE_MELEE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_Satchel: {
		ID:            WEAPON_Satchel,
		Name:          "Satchel Charge",
		Type:          WEAPON_TYPE_SATCHEL,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_MILLS_BOMB: {
		ID:            WEAPON_MILLS_BOMB,
		Name:          "Mills Bomb",
		Type:          WEAPON_TYPE_GRENADE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_NO_82_GRENADE: {
		ID:            WEAPON_NO_82_GRENADE,
		Name:          "Gammon Bomb",
		Type:          WEAPON_TYPE_GRENADE,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_PIAT: {
		ID:            WEAPON_PIAT,
		Name:          "PIAT",
		Type:          WEAPON_TYPE_ROCKET_LAUNCHER,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_BOYS_ANTI_TANK_RIFLE: {
		ID:            WEAPON_BOYS_ANTI_TANK_RIFLE,
		Name:          "Boys AT Rifle",
		Type:          WEAPON_TYPE_ANTI_MATERIEL_RIFLE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_A_P_SHRAPNEL_MINE_MK_II: {
		ID:            WEAPON_A_P_SHRAPNEL_MINE_MK_II,
		Name:          "AP Shrapnel Mine Mk II",
		Type:          WEAPON_TYPE_ANTI_PERSONNEL_MINE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_A_T_MINE_G_S_MK_V: {
		ID:            WEAPON_A_T_MINE_G_S_MK_V,
		Name:          "AT Mine G.S. Mk V",
		Type:          WEAPON_TYPE_ANTI_TANK_MINE,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_NO_2_MK_5_FLARE_PISTOL: {
		ID:            WEAPON_NO_2_MK_5_FLARE_PISTOL,
		Name:          "No.2 Mk V Flare Gun",
		Type:          WEAPON_TYPE_FLARE_GUN,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_6_POUNDER_QF_6_POUNDER: {
		ID:            WEAPON_QF_6_POUNDER_QF_6_POUNDER,
		Name:          "57mm Cannon",
		Type:          WEAPON_TYPE_ANTI_TANK_GUN,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_25_POUNDER_QF_25_POUNDER: {
		ID:            WEAPON_QF_25_POUNDER_QF_25_POUNDER,
		Name:          "88mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_DAIMLER: {
		ID:            WEAPON_DAIMLER,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_TETRARCH: {
		ID:            WEAPON_TETRARCH,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_M3_STUART_HONEY: {
		ID:            WEAPON_M3_STUART_HONEY,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_CROMWELL: {
		ID:            WEAPON_CROMWELL,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_CRUSADER_MK_III: {
		ID:            WEAPON_CRUSADER_MK_III,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_FIREFLY: {
		ID:            WEAPON_FIREFLY,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_CHURCHILL_MK_III: {
		ID:            WEAPON_CHURCHILL_MK_III,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_CHURCHILL_MK_VII: {
		ID:            WEAPON_CHURCHILL_MK_VII,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_BEDFORD_OYD_SUPPLY: {
		ID:            WEAPON_BEDFORD_OYD_SUPPLY,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_BEDFORD_OYD_TRANSPORT: {
		ID:            WEAPON_BEDFORD_OYD_TRANSPORT,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_CHURCHILL_MK_III_A_V_R_E: {
		ID:            WEAPON_CHURCHILL_MK_III_A_V_R_E,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_BISHOP_SP_25PDR: {
		ID:            WEAPON_BISHOP_SP_25PDR,
		Name:          "Roadkill",
		Type:          WEAPON_TYPE_ROADKILL,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_2_POUNDER_DAIMLER: {
		ID:            WEAPON_QF_2_POUNDER_DAIMLER,
		Name:          "QF 2-Pounder",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_DAIMLER: {
		ID:            WEAPON_COAXIAL_BESA_DAIMLER,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_2_POUNDER_TETRARCH: {
		ID:            WEAPON_QF_2_POUNDER_TETRARCH,
		Name:          "QF 2-Pounder",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_TETRARCH: {
		ID:            WEAPON_COAXIAL_BESA_TETRARCH,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_37MM_CANNON_M3_STUART_HONEY: {
		ID:            WEAPON_37MM_CANNON_M3_STUART_HONEY,
		Name:          "37mm Cannon",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_M3_STUART_HONEY: {
		ID:            WEAPON_COAXIAL_M1919_M3_STUART_HONEY,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_HULL_M1919_M3_STUART_HONEY: {
		ID:            WEAPON_HULL_M1919_M3_STUART_HONEY,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_OQF_75MM_CROMWELL: {
		ID:            WEAPON_OQF_75MM_CROMWELL,
		Name:          "QF 75mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_CROMWELL: {
		ID:            WEAPON_COAXIAL_BESA_CROMWELL,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_HULL_BESA_CROMWELL: {
		ID:            WEAPON_HULL_BESA_CROMWELL,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_OQF_57MM_CRUSADER_MK_III: {
		ID:            WEAPON_OQF_57MM_CRUSADER_MK_III,
		Name:          "QF 57mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_CRUSADER_MK_III: {
		ID:            WEAPON_COAXIAL_BESA_CRUSADER_MK_III,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_17_POUNDER_FIREFLY: {
		ID:            WEAPON_QF_17_POUNDER_FIREFLY,
		Name:          "QF 17-Pounder",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_COAXIAL_M1919_FIREFLY: {
		ID:            WEAPON_COAXIAL_M1919_FIREFLY,
		Name:          "M1919 Browning",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_OQF_57MM_CHURCHILL_MK_III: {
		ID:            WEAPON_OQF_57MM_CHURCHILL_MK_III,
		Name:          "QF 57mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III: {
		ID:            WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III: {
		ID:            WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_OQF_75MM_CHURCHILL_MK_VII: {
		ID:            WEAPON_OQF_75MM_CHURCHILL_MK_VII,
		Name:          "QF 75mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_VII: {
		ID:            WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_VII,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_VII: {
		ID:            WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_VII,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_230MM_PETARD_CHURCHILL_MK_III_A_V_R_E: {
		ID:            WEAPON_230MM_PETARD_CHURCHILL_MK_III_A_V_R_E,
		Name:          "230mm Petard",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E: {
		ID:            WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E: {
		ID:            WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_QF_25_POUNDER_BISHOP_SP_25PDR: {
		ID:            WEAPON_QF_25_POUNDER_BISHOP_SP_25PDR,
		Name:          "88mm Howitzer",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_6_POUNDER: {
		ID:            WEAPON_QF_6_POUNDER,
		Name:          "57mm Cannon",
		Type:          WEAPON_TYPE_ANTI_TANK_GUN,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_25_POUNDER: {
		ID:            WEAPON_QF_25_POUNDER,
		Name:          "88mm Howitzer",
		Type:          WEAPON_TYPE_ARTILLERY,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_2_POUNDER: {
		ID:            WEAPON_QF_2_POUNDER,
		Name:          "QF 2-Pounder",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_OQF_75MM: {
		ID:            WEAPON_OQF_75MM,
		Name:          "QF 75mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_OQF_57MM: {
		ID:            WEAPON_OQF_57MM,
		Name:          "QF 57mm",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_17_POUNDER: {
		ID:            WEAPON_QF_17_POUNDER,
		Name:          "QF 17-Pounder",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_CW},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA: {
		ID:            WEAPON_COAXIAL_BESA,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_COAXIAL_BESA_7_92MM: {
		ID:            WEAPON_COAXIAL_BESA_7_92MM,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_COAXIAL_MG,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_HULL_BESA: {
		ID:            WEAPON_HULL_BESA,
		Name:          "BESA",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_HULL_BESA_7_92MM: {
		ID:            WEAPON_HULL_BESA_7_92MM,
		Name:          "7.92mm",
		Type:          WEAPON_TYPE_TANK_HULL_MG,
		Factions:      []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_QF_25_POUNDER_GUN: {
		ID:            WEAPON_QF_25_POUNDER_GUN,
		Name:          "QF 25-Pounder",
		Type:          WEAPON_TYPE_TANK_CANNON,
		Factions:      []FactionIdentifier{FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_UNKNOWN: {
		ID:            WEAPON_UNKNOWN,
		Name:          "Unknown",
		Type:          WEAPON_TYPE_UNKNOWN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_SOV, FACTION_CW, FACTION_DAK, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_BOMBING_RUN: {
		ID:            WEAPON_BOMBING_RUN,
		Name:          "Bombing Run",
		Type:          WEAPON_TYPE_COMMANDER_ABILITY,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_CW, FACTION_DAK, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_STRAFING_RUN: {
		ID:            WEAPON_STRAFING_RUN,
		Name:          "Strafing Run",
		Type:          WEAPON_TYPE_COMMANDER_ABILITY,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_SOV, FACTION_CW, FACTION_DAK, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_PRECISION_STRIKE: {
		ID:            WEAPON_PRECISION_STRIKE,
		Name:          "Precision Strike",
		Type:          WEAPON_TYPE_COMMANDER_ABILITY,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_SOV, FACTION_CW, FACTION_DAK, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_Unknown: {
		ID:            WEAPON_Unknown,
		Name:          "Artillery Strike",
		Type:          WEAPON_TYPE_COMMANDER_ABILITY,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_SOV, FACTION_CW, FACTION_DAK, FACTION_B8A},
		Magnification: 0,
	},
	WEAPON_FIRESPOT: {
		ID:            WEAPON_FIRESPOT,
		Name:          "Fire",
		Type:          WEAPON_TYPE_UNKNOWN,
		Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_SOV, FACTION_CW, FACTION_DAK, FACTION_B8A},
		Magnification: 0,
	},
}

var fallback_weapon = Weapon{
	ID:            WEAPON_UNKNOWN,
	Name:          string(WEAPON_UNKNOWN),
	Type:          WEAPON_TYPE_UNKNOWN,
	Factions:      []FactionIdentifier{FACTION_GER, FACTION_US, FACTION_SOV, FACTION_CW, FACTION_DAK, FACTION_B8A},
	Magnification: 0,
}

func (w WeaponIdentifier) Weapon() Weapon {
	weapon, _ := ParseWeapon(string(w))
	return weapon
}

func ParseWeapon(weaponIdentifier string) (Weapon, error) {
	wi := WeaponIdentifier(weaponIdentifier)
	if weapon, ok := weaponMap[wi]; ok {
		return weapon, nil
	}
	for _, v := range weaponMap {
		if strings.HasPrefix(string(v.ID), weaponIdentifier) {
			logger.Debug("Using", v.ID, "as fallback for", weaponIdentifier)
			return v, nil
		}
	}
	return fallback_weapon, fmt.Errorf("weapon not found: %s", weaponIdentifier)
}

func AllWeapons() []Weapon {
	weapons := []Weapon{}
	for _, w := range weaponMap {
		weapons = append(weapons, w)
	}
	return weapons
}

func WeaponsByFaction(faction FactionIdentifier) []Weapon {
	weapons := []Weapon{}
	for _, w := range weaponMap {
		if slices.Contains(w.Factions, faction) {
			weapons = append(weapons, w)
		}
	}
	return weapons
}

func WeaponsByType(weaponType WeaponType) []Weapon {
	weapons := []Weapon{}
	for _, w := range weaponMap {
		if w.Type == weaponType {
			weapons = append(weapons, w)
		}
	}
	return weapons
}
