package hll

import (
	"strings"

	"github.com/zMoooooritz/go-let-loose/internal/logger"
)

type WeaponIdentifier string

const (
	WI_M1_GARAND                            WeaponIdentifier = "M1 GARAND"
	WI_M1_CARBINE                           WeaponIdentifier = "M1 CARBINE"
	WI_M1A1_THOMPSON                        WeaponIdentifier = "M1A1 THOMPSON"
	WI_M3_GREASE_GUN                        WeaponIdentifier = "M3 GREASE GUN"
	WI_M1918A2_BAR                          WeaponIdentifier = "M1918A2 BAR"
	WI_BROWNING_M1919                       WeaponIdentifier = "BROWNING M1919"
	WI_M1903_SPRINGFIELD                    WeaponIdentifier = "M1903 SPRINGFIELD"
	WI_M97_TRENCH_GUN                       WeaponIdentifier = "M97 TRENCH GUN"
	WI_COLT_M1911                           WeaponIdentifier = "COLT M1911"
	WI_M3_KNIFE                             WeaponIdentifier = "M3 KNIFE"
	WI_SATCHEL                              WeaponIdentifier = "SATCHEL"
	WI_MK2_GRENADE                          WeaponIdentifier = "MK2 GRENADE"
	WI_M2_FLAMETHROWER                      WeaponIdentifier = "M2 FLAMETHROWER"
	WI_BAZOOKA                              WeaponIdentifier = "BAZOOKA"
	WI_M2_AP_MINE                           WeaponIdentifier = "M2 AP MINE"
	WI_M1A1_AT_MINE                         WeaponIdentifier = "M1A1 AT MINE"
	WI_57MM_CANNON_M1_57mm                  WeaponIdentifier = "57MM CANNON [M1 57mm]"
	WI_155MM_HOWITZER_M114                  WeaponIdentifier = "155MM HOWITZER [M114]"
	WI_M8_Greyhound                         WeaponIdentifier = "M8 Greyhound"
	WI_Stuart_M5A1                          WeaponIdentifier = "Stuart M5A1"
	WI_Sherman_M4A1                         WeaponIdentifier = "Sherman M4A1"
	WI_Sherman_M4A3_75_W                    WeaponIdentifier = "Sherman M4A3(75)W"
	WI_Sherman_M4A3E2                       WeaponIdentifier = "Sherman M4A3E2"
	WI_Sherman_M4A3E2_76                    WeaponIdentifier = "Sherman M4A3E2(76)"
	WI_GMC_CCKW_353_Supply                  WeaponIdentifier = "GMC CCKW 353 (Supply)"
	WI_GMC_CCKW_353_Transport               WeaponIdentifier = "GMC CCKW 353 (Transport)"
	WI_M3_Half_track                        WeaponIdentifier = "M3 Half-track"
	WI_Jeep_Willys                          WeaponIdentifier = "Jeep Willys"
	WI_M6_37mm_M8_Greyhound                 WeaponIdentifier = "M6 37mm [M8 Greyhound]"
	WI_COAXIAL_M1919_M8_Greyhound           WeaponIdentifier = "COAXIAL M1919 [M8 Greyhound]"
	WI_37MM_CANNON_Stuart_M5A1              WeaponIdentifier = "37MM CANNON [Stuart M5A1]"
	WI_COAXIAL_M1919_Stuart_M5A1            WeaponIdentifier = "COAXIAL M1919 [Stuart M5A1]"
	WI_HULL_M1919_Stuart_M5A1               WeaponIdentifier = "HULL M1919 [Stuart M5A1]"
	WI_75MM_CANNON_Sherman_M4A1             WeaponIdentifier = "75MM CANNON [Sherman M4A1]"
	WI_COAXIAL_M1919_Sherman_M4A1           WeaponIdentifier = "COAXIAL M1919 [Sherman M4A1]"
	WI_HULL_M1919_Sherman_M4A1              WeaponIdentifier = "HULL M1919 [Sherman M4A1]"
	WI_75MM_CANNON_Sherman_M4A3_75_W        WeaponIdentifier = "75MM CANNON [Sherman M4A3(75)W]"
	WI_COAXIAL_M1919_Sherman_M4A3_75_W      WeaponIdentifier = "COAXIAL M1919 [Sherman M4A3(75)W]"
	WI_HULL_M1919_Sherman_M4A3_75_W         WeaponIdentifier = "HULL M1919 [Sherman M4A3(75)W]"
	WI_75MM_M3_GUN_Sherman_M4A3E2           WeaponIdentifier = "75MM M3 GUN [Sherman M4A3E2]"
	WI_COAXIAL_M1919_Sherman_M4A3E2         WeaponIdentifier = "COAXIAL M1919 [Sherman M4A3E2]"
	WI_HULL_M1919_Sherman_M4A3E2            WeaponIdentifier = "HULL M1919 [Sherman M4A3E2]"
	WI_76MM_M1_GUN_Sherman_M4A3E2_76        WeaponIdentifier = "76MM M1 GUN [Sherman M4A3E2(76)]"
	WI_COAXIAL_M1919_Sherman_M4A3E2_76      WeaponIdentifier = "COAXIAL M1919 [Sherman M4A3E2(76)]"
	WI_HULL_M1919_Sherman_M4A3E2_76         WeaponIdentifier = "HULL M1919 [Sherman M4A3E2(76)]"
	WI_M2_Browning_M3_Half_track            WeaponIdentifier = "M2 Browning [M3 Half-track]"
	WI_KARABINER_98K                        WeaponIdentifier = "KARABINER 98K"
	WI_GEWEHR_43                            WeaponIdentifier = "GEWEHR 43"
	WI_STG44                                WeaponIdentifier = "STG44"
	WI_FG42                                 WeaponIdentifier = "FG42"
	WI_MP40                                 WeaponIdentifier = "MP40"
	WI_MG34                                 WeaponIdentifier = "MG34"
	WI_MG42                                 WeaponIdentifier = "MG42"
	WI_FLAMMENWERFER_41                     WeaponIdentifier = "FLAMMENWERFER 41"
	WI_KARABINER_98K_x8                     WeaponIdentifier = "KARABINER 98K x8"
	WI_FG42_x4                              WeaponIdentifier = "FG42 x4"
	WI_LUGER_P08                            WeaponIdentifier = "LUGER P08"
	WI_WALTHER_P38                          WeaponIdentifier = "WALTHER P38"
	WI_FELDSPATEN                           WeaponIdentifier = "FELDSPATEN"
	WI_M24_STIELHANDGRANATE                 WeaponIdentifier = "M24 STIELHANDGRANATE"
	WI_M43_STIELHANDGRANATE                 WeaponIdentifier = "M43 STIELHANDGRANATE"
	WI_PANZERSCHRECK                        WeaponIdentifier = "PANZERSCHRECK"
	WI_S_MINE                               WeaponIdentifier = "S-MINE"
	WI_TELLERMINE_43                        WeaponIdentifier = "TELLERMINE 43"
	WI_75MM_CANNON_PAK_40                   WeaponIdentifier = "75MM CANNON [PAK 40]"
	WI_150MM_HOWITZER_sFH_18                WeaponIdentifier = "150MM HOWITZER [sFH 18]"
	WI_Sd_Kfz_234_Puma                      WeaponIdentifier = "Sd.Kfz.234 Puma"
	WI_Sd_Kfz_121_Luchs                     WeaponIdentifier = "Sd.Kfz.121 Luchs"
	WI_Sd_Kfz_161_Panzer_IV                 WeaponIdentifier = "Sd.Kfz.161 Panzer IV"
	WI_Sd_Kfz_171_Panther                   WeaponIdentifier = "Sd.Kfz.171 Panther"
	WI_Sd_Kfz_181_Tiger_1                   WeaponIdentifier = "Sd.Kfz.181 Tiger 1"
	WI_Opel_Blitz_Supply                    WeaponIdentifier = "Opel Blitz (Supply)"
	WI_Opel_Blitz_Transport                 WeaponIdentifier = "Opel Blitz (Transport)"
	WI_Sd_Kfz_251_Half_track                WeaponIdentifier = "Sd.Kfz 251 Half-track"
	WI_Kubelwagen                           WeaponIdentifier = "Kubelwagen"
	WI_50mm_KwK_39_1_Sd_Kfz_234_Puma        WeaponIdentifier = "50mm KwK 39/1 [Sd.Kfz.234 Puma]"
	WI_COAXIAL_MG34_Sd_Kfz_234_Puma         WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.234 Puma]"
	WI_20MM_KWK_30_Sd_Kfz_121_Luchs         WeaponIdentifier = "20MM KWK 30 [Sd.Kfz.121 Luchs]"
	WI_COAXIAL_MG34_Sd_Kfz_121_Luchs        WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.121 Luchs]"
	WI_75MM_CANNON_Sd_Kfz_161_Panzer_IV     WeaponIdentifier = "75MM CANNON [Sd.Kfz.161 Panzer IV]"
	WI_COAXIAL_MG34_Sd_Kfz_161_Panzer_IV    WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.161 Panzer IV]"
	WI_HULL_MG34_Sd_Kfz_161_Panzer_IV       WeaponIdentifier = "HULL MG34 [Sd.Kfz.161 Panzer IV]"
	WI_75MM_CANNON_Sd_Kfz_171_Panther       WeaponIdentifier = "75MM CANNON [Sd.Kfz.171 Panther]"
	WI_COAXIAL_MG34_Sd_Kfz_171_Panther      WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.171 Panther]"
	WI_HULL_MG34_Sd_Kfz_171_Panther         WeaponIdentifier = "HULL MG34 [Sd.Kfz.171 Panther]"
	WI_88_KWK_36_L_56_Sd_Kfz_181_Tiger_1    WeaponIdentifier = "88 KWK 36 L/56 [Sd.Kfz.181 Tiger 1]"
	WI_COAXIAL_MG34_Sd_Kfz_181_Tiger_1      WeaponIdentifier = "COAXIAL MG34 [Sd.Kfz.181 Tiger 1]"
	WI_HULL_MG34_Sd_Kfz_181_Tiger_1         WeaponIdentifier = "HULL MG34 [Sd.Kfz.181 Tiger 1]"
	WI_MG_42_Sd_Kfz_251_Half_track          WeaponIdentifier = "MG 42 [Sd.Kfz 251 Half-track]"
	WI_MOSIN_NAGANT_1891                    WeaponIdentifier = "MOSIN NAGANT 1891"
	WI_MOSIN_NAGANT_91_30                   WeaponIdentifier = "MOSIN NAGANT 91/30"
	WI_MOSIN_NAGANT_M38                     WeaponIdentifier = "MOSIN NAGANT M38"
	WI_SVT40                                WeaponIdentifier = "SVT40"
	WI_PPSH_41                              WeaponIdentifier = "PPSH 41"
	WI_PPSH_41_W_DRUM                       WeaponIdentifier = "PPSH 41 W/DRUM"
	WI_DP_27                                WeaponIdentifier = "DP-27"
	WI_SCOPED_MOSIN_NAGANT_91_30            WeaponIdentifier = "SCOPED MOSIN NAGANT 91/30"
	WI_SCOPED_SVT40                         WeaponIdentifier = "SCOPED SVT40"
	WI_NAGANT_M1895                         WeaponIdentifier = "NAGANT M1895"
	WI_TOKAREV_TT33                         WeaponIdentifier = "TOKAREV TT33"
	WI_MPL_50_SPADE                         WeaponIdentifier = "MPL-50 SPADE"
	WI_SATCHEL_CHARGE                       WeaponIdentifier = "SATCHEL CHARGE"
	WI_RG_42_GRENADE                        WeaponIdentifier = "RG-42 GRENADE"
	WI_MOLOTOV                              WeaponIdentifier = "MOLOTOV"
	WI_PTRS_41                              WeaponIdentifier = "PTRS-41"
	WI_POMZ_AP_MINE                         WeaponIdentifier = "POMZ AP MINE"
	WI_TM_35_AT_MINE                        WeaponIdentifier = "TM-35 AT MINE"
	WI_57MM_CANNON_ZiS_2                    WeaponIdentifier = "57MM CANNON [ZiS-2]"
	WI_122MM_HOWITZER_M1938_M_30            WeaponIdentifier = "122MM HOWITZER [M1938 (M-30)]"
	WI_BA_10                                WeaponIdentifier = "BA-10"
	WI_T70                                  WeaponIdentifier = "T70"
	WI_T34_76                               WeaponIdentifier = "T34/76"
	WI_IS_1                                 WeaponIdentifier = "IS-1"
	WI_ZIS_5_Supply                         WeaponIdentifier = "ZIS-5 (Supply)"
	WI_ZIS_5_Transport                      WeaponIdentifier = "ZIS-5 (Transport)"
	WI_GAZ_67                               WeaponIdentifier = "GAZ-67"
	WI_19_K_45MM_BA_10                      WeaponIdentifier = "19-K 45MM [BA-10]"
	WI_COAXIAL_DT_BA_10                     WeaponIdentifier = "COAXIAL DT [BA-10]"
	WI_45MM_M1937_T70                       WeaponIdentifier = "45MM M1937 [T70]"
	WI_COAXIAL_DT_T70                       WeaponIdentifier = "COAXIAL DT [T70]"
	WI_76MM_ZiS_5_T34_76                    WeaponIdentifier = "76MM ZiS-5 [T34/76]"
	WI_COAXIAL_DT_T34_76                    WeaponIdentifier = "COAXIAL DT [T34/76]"
	WI_HULL_DT_T34_76                       WeaponIdentifier = "HULL DT [T34/76]"
	WI_D_5T_85MM_IS_1                       WeaponIdentifier = "D-5T 85MM [IS-1]"
	WI_COAXIAL_DT_IS_1                      WeaponIdentifier = "COAXIAL DT [IS-1]"
	WI_HULL_DT_IS_1                         WeaponIdentifier = "HULL DT [IS-1]"
	WI_SMLE_No_1_Mk_III                     WeaponIdentifier = "SMLE No.1 Mk III"
	WI_Rifle_No_5_Mk_I                      WeaponIdentifier = "Rifle No.5 Mk I"
	WI_Rifle_No_4_Mk_I                      WeaponIdentifier = "Rifle No.4 Mk I"
	WI_Sten_Gun_Mk_II                       WeaponIdentifier = "Sten Gun Mk.II"
	WI_Sten_Gun_Mk_V                        WeaponIdentifier = "Sten Gun Mk.V"
	WI_Lanchester                           WeaponIdentifier = "Lanchester"
	WI_M1928A1_THOMPSON                     WeaponIdentifier = "M1928A1 THOMPSON"
	WI_Bren_Gun                             WeaponIdentifier = "Bren Gun"
	WI_Lewis_Gun                            WeaponIdentifier = "Lewis Gun"
	WI_FLAMETHROWER                         WeaponIdentifier = "FLAMETHROWER"
	WI_Lee_Enfield_Pattern_1914_Sniper      WeaponIdentifier = "Lee-Enfield Pattern 1914 Sniper"
	WI_Rifle_No_4_Mk_I_Sniper               WeaponIdentifier = "Rifle No.4 Mk I Sniper"
	WI_Webley_MK_VI                         WeaponIdentifier = "Webley MK VI"
	WI_Fairbairn_Sykes                      WeaponIdentifier = "Fairbairn–Sykes"
	WI_Satchel                              WeaponIdentifier = "Satchel"
	WI_Mills_Bomb                           WeaponIdentifier = "Mills Bomb"
	WI_NO_82_Grenade                        WeaponIdentifier = "No.82 Grenade"
	WI_PIAT                                 WeaponIdentifier = "PIAT"
	WI_Boys_Anti_tank_Rifle                 WeaponIdentifier = "Boys Anti-tank Rifle"
	WI_A_P_Shrapnel_Mine_Mk_II              WeaponIdentifier = "A.P. Shrapnel Mine Mk II"
	WI_A_T_Mine_G_S_Mk_V                    WeaponIdentifier = "A.T. Mine G.S. Mk V"
	WI_QF_6_POUNDER_QF_6_Pounder            WeaponIdentifier = "QF 6-POUNDER [QF 6-Pounder]"
	WI_QF_25_POUNDER_QF_25_Pounder          WeaponIdentifier = "QF 25-POUNDER [QF 25-Pounder]"
	WI_Daimler                              WeaponIdentifier = "Daimler"
	WI_Tetrarch                             WeaponIdentifier = "Tetrarch"
	WI_M3_Stuart_Honey                      WeaponIdentifier = "M3 Stuart Honey"
	WI_Cromwell                             WeaponIdentifier = "Cromwell"
	WI_Crusader_Mk_III                      WeaponIdentifier = "Crusader Mk.III"
	WI_Firefly                              WeaponIdentifier = "Firefly"
	WI_Churchill_Mk_III                     WeaponIdentifier = "Churchill Mk.III"
	WI_Churchill_Mk_VII                     WeaponIdentifier = "Churchill Mk.VII"
	WI_Bedford_OYD_Supply                   WeaponIdentifier = "Bedford OYD (Supply)"
	WI_Bedford_OYD_Transport                WeaponIdentifier = "Bedford OYD (Transport)"
	WI_QF_2_POUNDER_Daimler                 WeaponIdentifier = "QF 2-POUNDER [Daimler]"
	WI_COAXIAL_BESA_Daimler                 WeaponIdentifier = "COAXIAL BESA [Daimler]"
	WI_QF_2_POUNDER_Tetrarch                WeaponIdentifier = "QF 2-POUNDER [Tetrarch]"
	WI_COAXIAL_BESA_Tetrarch                WeaponIdentifier = "COAXIAL BESA [Tetrarch]"
	WI_37MM_CANNON_M3_Stuart_Honey          WeaponIdentifier = "37MM CANNON [M3 Stuart Honey]"
	WI_COAXIAL_M1919_M3_Stuart_Honey        WeaponIdentifier = "COAXIAL M1919 [M3 Stuart Honey]"
	WI_HULL_M1919_M3_Stuart_Honey           WeaponIdentifier = "HULL M1919 [M3 Stuart Honey]"
	WI_OQF_75MM_Cromwell                    WeaponIdentifier = "OQF 75MM [Cromwell]"
	WI_COAXIAL_BESA_Cromwell                WeaponIdentifier = "COAXIAL BESA [Cromwell]"
	WI_HULL_BESA_Cromwell                   WeaponIdentifier = "HULL BESA [Cromwell]"
	WI_OQF_57MM_Crusader_Mk_III             WeaponIdentifier = "OQF 57MM [Crusader Mk.III]"
	WI_COAXIAL_BESA_Crusader_Mk_III         WeaponIdentifier = "COAXIAL BESA [Crusader Mk.III]"
	WI_QF_17_POUNDER_Firefly                WeaponIdentifier = "QF 17-POUNDER [Firefly]"
	WI_COAXIAL_M1919_Firefly                WeaponIdentifier = "COAXIAL M1919 [Firefly]"
	WI_OQF_57MM_Churchill_Mk_III            WeaponIdentifier = "OQF 57MM [Churchill Mk.III]"
	WI_COAXIAL_BESA_7_92mm_Churchill_Mk_III WeaponIdentifier = "COAXIAL BESA 7.92mm [Churchill Mk.III]"
	WI_HULL_BESA_7_92mm_Churchill_Mk_III    WeaponIdentifier = "HULL BESA 7.92mm [Churchill Mk.III]"
	WI_OQF_75MM_Churchill_Mk_VII            WeaponIdentifier = "OQF 75MM [Churchill Mk.VII]"
	WI_COAXIAL_BESA_7_92mm_Churchill_Mk_VII WeaponIdentifier = "COAXIAL BESA 7.92mm [Churchill Mk.VII]"
	WI_HULL_BESA_7_92mm_Churchill_Mk_VII    WeaponIdentifier = "HULL BESA 7.92mm [Churchill Mk.VII]"
	WI_UNKNOWN                              WeaponIdentifier = "UNKNOWN"
	WI_BOMBING_RUN                          WeaponIdentifier = "BOMBING RUN"
	WI_STRAFING_RUN                         WeaponIdentifier = "STRAFING RUN"
	WI_PRECISION_STRIKE                     WeaponIdentifier = "PRECISION STRIKE"
	WI_Unknown                              WeaponIdentifier = "Unknown"
	WI_FLARE_GUN                            WeaponIdentifier = "FLARE GUN"
	WI_INVALID                              WeaponIdentifier = "INVALID"
)

type WeaponCategory int

const (
	WcSubmachineGun WeaponCategory = iota
	WcSemiAutoRifle
	WcBoltActionRifle
	WcAssaultRifle
	WcShotgun
	WcMachinegun
	WcSniperRifle
	WcPistol
	WcFlamethrower
	WcMeele
	WcGrenade
	WcSatchel
	WcAntiPersonnelMine
	WcAntiTankMine
	WcAntiTankRifle
	WcFlaregun
	WcArtillerygun
	WcAntiTankGun
	WcVehicle
	WcMainCannon
	WcCoaxialMachinegun
	WcHullMachinegun
	WcMountedMachinegun
	WcCommanderAbility
	WcUnknown
)

type Weapon struct {
	ID       WeaponIdentifier
	Name     string
	Factions []Faction
	Category WeaponCategory
}

var weapons = map[WeaponIdentifier]Weapon{
	WI_M1_GARAND:                       {ID: WI_M1_GARAND, Name: "M1 Garand", Factions: []Faction{FctUS}, Category: WcSemiAutoRifle},
	WI_M1_CARBINE:                      {ID: WI_M1_CARBINE, Name: "M1 Carbine", Factions: []Faction{FctUS}, Category: WcSemiAutoRifle},
	WI_M1A1_THOMPSON:                   {ID: WI_M1A1_THOMPSON, Name: "M1A1 Thompson", Factions: []Faction{FctUS}, Category: WcSubmachineGun},
	WI_M3_GREASE_GUN:                   {ID: WI_M3_GREASE_GUN, Name: "M3 Grease Gun", Factions: []Faction{FctUS}, Category: WcSubmachineGun},
	WI_M1918A2_BAR:                     {ID: WI_M1918A2_BAR, Name: "M1918A2 BAR", Factions: []Faction{FctUS}, Category: WcAssaultRifle},
	WI_BROWNING_M1919:                  {ID: WI_BROWNING_M1919, Name: "M1919 Browning", Factions: []Faction{FctUS}, Category: WcMachinegun},
	WI_M1903_SPRINGFIELD:               {ID: WI_M1903_SPRINGFIELD, Name: "M1903 Springfield (4x)", Factions: []Faction{FctUS}, Category: WcSniperRifle},
	WI_M97_TRENCH_GUN:                  {ID: WI_M97_TRENCH_GUN, Name: "M97 Trench Gun", Factions: []Faction{FctUS}, Category: WcShotgun},
	WI_COLT_M1911:                      {ID: WI_COLT_M1911, Name: "Colt M1911", Factions: []Faction{FctUS}, Category: WcPistol},
	WI_M3_KNIFE:                        {ID: WI_M3_KNIFE, Name: "US Melee", Factions: []Faction{FctUS}, Category: WcMeele},
	WI_SATCHEL:                         {ID: WI_SATCHEL, Name: "Satchel Charge", Factions: []Faction{FctUS}, Category: WcSatchel},
	WI_MK2_GRENADE:                     {ID: WI_MK2_GRENADE, Name: "US Grenade", Factions: []Faction{FctUS}, Category: WcGrenade},
	WI_M2_FLAMETHROWER:                 {ID: WI_M2_FLAMETHROWER, Name: "US Flamethrower", Factions: []Faction{FctUS}, Category: WcFlamethrower},
	WI_BAZOOKA:                         {ID: WI_BAZOOKA, Name: "Bazooka", Factions: []Faction{FctUS, FctRUS}, Category: WcAntiTankRifle},
	WI_M2_AP_MINE:                      {ID: WI_M2_AP_MINE, Name: "US AP Mine", Factions: []Faction{FctUS}, Category: WcAntiPersonnelMine},
	WI_M1A1_AT_MINE:                    {ID: WI_M1A1_AT_MINE, Name: "US AT Mine", Factions: []Faction{FctUS}, Category: WcAntiTankMine},
	WI_57MM_CANNON_M1_57mm:             {ID: WI_57MM_CANNON_M1_57mm, Name: "US AT Gun", Factions: []Faction{FctUS}, Category: WcAntiTankGun},
	WI_155MM_HOWITZER_M114:             {ID: WI_155MM_HOWITZER_M114, Name: "US Artillery", Factions: []Faction{FctUS}, Category: WcArtillerygun},
	WI_M8_Greyhound:                    {ID: WI_M8_Greyhound, Name: "US Roadkill [M8 Greyhound]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_Stuart_M5A1:                     {ID: WI_Stuart_M5A1, Name: "US Roadkill [Stuart M5A1]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_Sherman_M4A1:                    {ID: WI_Sherman_M4A1, Name: "US Roadkill [Sherman M4]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_Sherman_M4A3_75_W:               {ID: WI_Sherman_M4A3_75_W, Name: "US Roadkill [Sherman M4A3 75w]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_Sherman_M4A3E2:                  {ID: WI_Sherman_M4A3E2, Name: "US Roadkill [Sherman 75mm]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_Sherman_M4A3E2_76:               {ID: WI_Sherman_M4A3E2_76, Name: "US Roadkill [Sherman 76mm]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_GMC_CCKW_353_Supply:             {ID: WI_GMC_CCKW_353_Supply, Name: "US Roadkill [US Supply Truck]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_GMC_CCKW_353_Transport:          {ID: WI_GMC_CCKW_353_Transport, Name: "US Roadkill [US Transport Truck]", Factions: []Faction{FctUS}, Category: WcVehicle},
	WI_M3_Half_track:                   {ID: WI_M3_Half_track, Name: "US Roadkill [US Half-track]", Factions: []Faction{FctUS, FctRUS, FctGB}, Category: WcVehicle},
	WI_Jeep_Willys:                     {ID: WI_Jeep_Willys, Name: "US Roadkill [US Jeep]", Factions: []Faction{FctUS, FctGB}, Category: WcVehicle},
	WI_M6_37mm_M8_Greyhound:            {ID: WI_M6_37mm_M8_Greyhound, Name: "US Tank Cannon [M8 Greyhound]", Factions: []Faction{FctUS}, Category: WcMainCannon},
	WI_COAXIAL_M1919_M8_Greyhound:      {ID: WI_COAXIAL_M1919_M8_Greyhound, Name: "US Tank Coaxial [M8 Greyhound]", Factions: []Faction{FctUS}, Category: WcCoaxialMachinegun},
	WI_37MM_CANNON_Stuart_M5A1:         {ID: WI_37MM_CANNON_Stuart_M5A1, Name: "US Tank Cannon [Stuart M5A1]", Factions: []Faction{FctUS}, Category: WcMainCannon},
	WI_COAXIAL_M1919_Stuart_M5A1:       {ID: WI_COAXIAL_M1919_Stuart_M5A1, Name: "US Tank Coaxial [Stuart M5A1]", Factions: []Faction{FctUS}, Category: WcCoaxialMachinegun},
	WI_HULL_M1919_Stuart_M5A1:          {ID: WI_HULL_M1919_Stuart_M5A1, Name: "US Tank Hull MG [Stuart M5A1]", Factions: []Faction{FctUS}, Category: WcHullMachinegun},
	WI_75MM_CANNON_Sherman_M4A1:        {ID: WI_75MM_CANNON_Sherman_M4A1, Name: "US Tank Cannon [Sherman M4]", Factions: []Faction{FctUS}, Category: WcMainCannon},
	WI_COAXIAL_M1919_Sherman_M4A1:      {ID: WI_COAXIAL_M1919_Sherman_M4A1, Name: "US Tank Coaxial [Sherman M4]", Factions: []Faction{FctUS}, Category: WcCoaxialMachinegun},
	WI_HULL_M1919_Sherman_M4A1:         {ID: WI_HULL_M1919_Sherman_M4A1, Name: "US Tank Hull MG [Sherman M4]", Factions: []Faction{FctUS}, Category: WcHullMachinegun},
	WI_75MM_CANNON_Sherman_M4A3_75_W:   {ID: WI_75MM_CANNON_Sherman_M4A3_75_W, Name: "US Tank Cannon [Sherman M4A3 75w]", Factions: []Faction{FctUS}, Category: WcMainCannon},
	WI_COAXIAL_M1919_Sherman_M4A3_75_W: {ID: WI_COAXIAL_M1919_Sherman_M4A3_75_W, Name: "US Tank Coaxial [Sherman M4A3 75w]", Factions: []Faction{FctUS}, Category: WcCoaxialMachinegun},
	WI_HULL_M1919_Sherman_M4A3_75_W:    {ID: WI_HULL_M1919_Sherman_M4A3_75_W, Name: "US Tank Hull MG [Sherman M4A3 75w]", Factions: []Faction{FctUS}, Category: WcHullMachinegun},
	WI_75MM_M3_GUN_Sherman_M4A3E2:      {ID: WI_75MM_M3_GUN_Sherman_M4A3E2, Name: "US Tank Cannon [Sherman 75mm]", Factions: []Faction{FctUS}, Category: WcMainCannon},
	WI_COAXIAL_M1919_Sherman_M4A3E2:    {ID: WI_COAXIAL_M1919_Sherman_M4A3E2, Name: "US Tank Coaxial [Sherman 75mm]", Factions: []Faction{FctUS}, Category: WcCoaxialMachinegun},
	WI_HULL_M1919_Sherman_M4A3E2:       {ID: WI_HULL_M1919_Sherman_M4A3E2, Name: "US Tank Hull MG [Sherman 75mm]", Factions: []Faction{FctUS}, Category: WcHullMachinegun},
	WI_76MM_M1_GUN_Sherman_M4A3E2_76:   {ID: WI_76MM_M1_GUN_Sherman_M4A3E2_76, Name: "US Tank Cannon [Sherman 76mm]", Factions: []Faction{FctUS}, Category: WcMainCannon},
	WI_COAXIAL_M1919_Sherman_M4A3E2_76: {ID: WI_COAXIAL_M1919_Sherman_M4A3E2_76, Name: "US Tank Coaxial [Sherman 76mm]", Factions: []Faction{FctUS}, Category: WcCoaxialMachinegun},
	WI_HULL_M1919_Sherman_M4A3E2_76:    {ID: WI_HULL_M1919_Sherman_M4A3E2_76, Name: "US Tank Hull MG [Sherman 76mm]", Factions: []Faction{FctUS}, Category: WcHullMachinegun},
	WI_M2_Browning_M3_Half_track:       {ID: WI_M2_Browning_M3_Half_track, Name: "US Half-track MG [US Half-track]", Factions: []Faction{FctUS, FctRUS, FctGB}, Category: WcMountedMachinegun},

	WI_KARABINER_98K:                     {ID: WI_KARABINER_98K, Name: "Kar98k", Factions: []Faction{FctGER}, Category: WcBoltActionRifle},
	WI_GEWEHR_43:                         {ID: WI_GEWEHR_43, Name: "G43", Factions: []Faction{FctGER}, Category: WcSemiAutoRifle},
	WI_STG44:                             {ID: WI_STG44, Name: "STG44", Factions: []Faction{FctGER}, Category: WcAssaultRifle},
	WI_FG42:                              {ID: WI_FG42, Name: "FG42", Factions: []Faction{FctGER}, Category: WcAssaultRifle},
	WI_MP40:                              {ID: WI_MP40, Name: "MP40", Factions: []Faction{FctGER}, Category: WcSubmachineGun},
	WI_MG34:                              {ID: WI_MG34, Name: "MG34", Factions: []Faction{FctGER}, Category: WcMachinegun},
	WI_MG42:                              {ID: WI_MG42, Name: "MG42", Factions: []Faction{FctGER}, Category: WcMachinegun},
	WI_FLAMMENWERFER_41:                  {ID: WI_FLAMMENWERFER_41, Name: "GER Flamethrower", Factions: []Faction{FctGER}, Category: WcFlamethrower},
	WI_KARABINER_98K_x8:                  {ID: WI_KARABINER_98K_x8, Name: "Kar98k (8x)", Factions: []Faction{FctGER}, Category: WcSniperRifle},
	WI_FG42_x4:                           {ID: WI_FG42_x4, Name: "FG42 (4x)", Factions: []Faction{FctGER}, Category: WcSniperRifle},
	WI_LUGER_P08:                         {ID: WI_LUGER_P08, Name: "Luger P08", Factions: []Faction{FctGER}, Category: WcPistol},
	WI_WALTHER_P38:                       {ID: WI_WALTHER_P38, Name: "Walther P38", Factions: []Faction{FctGER}, Category: WcPistol},
	WI_FELDSPATEN:                        {ID: WI_FELDSPATEN, Name: "GER Melee", Factions: []Faction{FctGER}, Category: WcMeele},
	WI_M24_STIELHANDGRANATE:              {ID: WI_M24_STIELHANDGRANATE, Name: "GER Grenade", Factions: []Faction{FctGER}, Category: WcGrenade},
	WI_M43_STIELHANDGRANATE:              {ID: WI_M43_STIELHANDGRANATE, Name: "GER Grenade", Factions: []Faction{FctGER}, Category: WcGrenade},
	WI_PANZERSCHRECK:                     {ID: WI_PANZERSCHRECK, Name: "Panzerschreck", Factions: []Faction{FctGER}, Category: WcAntiTankRifle},
	WI_S_MINE:                            {ID: WI_S_MINE, Name: "GER AP Mine", Factions: []Faction{FctGER}, Category: WcAntiPersonnelMine},
	WI_TELLERMINE_43:                     {ID: WI_TELLERMINE_43, Name: "GER AT Mine", Factions: []Faction{FctGER}, Category: WcAntiTankMine},
	WI_75MM_CANNON_PAK_40:                {ID: WI_75MM_CANNON_PAK_40, Name: "GER AT Gun", Factions: []Faction{FctGER}, Category: WcAntiTankGun},
	WI_150MM_HOWITZER_sFH_18:             {ID: WI_150MM_HOWITZER_sFH_18, Name: "GER Artillery", Factions: []Faction{FctGER}, Category: WcArtillerygun},
	WI_Sd_Kfz_234_Puma:                   {ID: WI_Sd_Kfz_234_Puma, Name: "GER Roadkill [Puma]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Sd_Kfz_121_Luchs:                  {ID: WI_Sd_Kfz_121_Luchs, Name: "GER Roadkill [Luchs]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Sd_Kfz_161_Panzer_IV:              {ID: WI_Sd_Kfz_161_Panzer_IV, Name: "GER Roadkill [Panzer IV]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Sd_Kfz_171_Panther:                {ID: WI_Sd_Kfz_171_Panther, Name: "GER Roadkill [Panther]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Sd_Kfz_181_Tiger_1:                {ID: WI_Sd_Kfz_181_Tiger_1, Name: "GER Roadkill [Tiger 1]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Opel_Blitz_Supply:                 {ID: WI_Opel_Blitz_Supply, Name: "GER Roadkill [GER Supply Truck]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Opel_Blitz_Transport:              {ID: WI_Opel_Blitz_Transport, Name: "GER Roadkill [GER Transport Truck]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Sd_Kfz_251_Half_track:             {ID: WI_Sd_Kfz_251_Half_track, Name: "GER Roadkill [GER Half-track]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_Kubelwagen:                        {ID: WI_Kubelwagen, Name: "GER Roadkill [GER Jeep]", Factions: []Faction{FctGER}, Category: WcVehicle},
	WI_50mm_KwK_39_1_Sd_Kfz_234_Puma:     {ID: WI_50mm_KwK_39_1_Sd_Kfz_234_Puma, Name: "GER Tank Cannon [Puma]", Factions: []Faction{FctGER}, Category: WcMainCannon},
	WI_COAXIAL_MG34_Sd_Kfz_234_Puma:      {ID: WI_COAXIAL_MG34_Sd_Kfz_234_Puma, Name: "GER Tank Coaxial [Puma]", Factions: []Faction{FctGER}, Category: WcCoaxialMachinegun},
	WI_20MM_KWK_30_Sd_Kfz_121_Luchs:      {ID: WI_20MM_KWK_30_Sd_Kfz_121_Luchs, Name: "GER Tank Cannon [Luchs]", Factions: []Faction{FctGER}, Category: WcMainCannon},
	WI_COAXIAL_MG34_Sd_Kfz_121_Luchs:     {ID: WI_COAXIAL_MG34_Sd_Kfz_121_Luchs, Name: "GER Tank Coaxial [Luchs]", Factions: []Faction{FctGER}, Category: WcCoaxialMachinegun},
	WI_75MM_CANNON_Sd_Kfz_161_Panzer_IV:  {ID: WI_75MM_CANNON_Sd_Kfz_161_Panzer_IV, Name: "GER Tank Cannon [Panzer IV]", Factions: []Faction{FctGER}, Category: WcMainCannon},
	WI_COAXIAL_MG34_Sd_Kfz_161_Panzer_IV: {ID: WI_COAXIAL_MG34_Sd_Kfz_161_Panzer_IV, Name: "GER Tank Coaxial [Panzer IV]", Factions: []Faction{FctGER}, Category: WcCoaxialMachinegun},
	WI_HULL_MG34_Sd_Kfz_161_Panzer_IV:    {ID: WI_HULL_MG34_Sd_Kfz_161_Panzer_IV, Name: "GER Tank Hull MG [Panzer IV]", Factions: []Faction{FctGER}, Category: WcHullMachinegun},
	WI_75MM_CANNON_Sd_Kfz_171_Panther:    {ID: WI_75MM_CANNON_Sd_Kfz_171_Panther, Name: "GER Tank Cannon [Panther]", Factions: []Faction{FctGER}, Category: WcMainCannon},
	WI_COAXIAL_MG34_Sd_Kfz_171_Panther:   {ID: WI_COAXIAL_MG34_Sd_Kfz_171_Panther, Name: "GER Tank Coaxial [Panther]", Factions: []Faction{FctGER}, Category: WcCoaxialMachinegun},
	WI_HULL_MG34_Sd_Kfz_171_Panther:      {ID: WI_HULL_MG34_Sd_Kfz_171_Panther, Name: "GER Tank Hull MG [Panther]", Factions: []Faction{FctGER}, Category: WcHullMachinegun},
	WI_88_KWK_36_L_56_Sd_Kfz_181_Tiger_1: {ID: WI_88_KWK_36_L_56_Sd_Kfz_181_Tiger_1, Name: "GER Tank Cannon [Tiger 1]", Factions: []Faction{FctGER}, Category: WcMainCannon},
	WI_COAXIAL_MG34_Sd_Kfz_181_Tiger_1:   {ID: WI_COAXIAL_MG34_Sd_Kfz_181_Tiger_1, Name: "GER Tank Coaxial [Tiger 1]", Factions: []Faction{FctGER}, Category: WcCoaxialMachinegun},
	WI_HULL_MG34_Sd_Kfz_181_Tiger_1:      {ID: WI_HULL_MG34_Sd_Kfz_181_Tiger_1, Name: "GER Tank Hull MG [Tiger 1]", Factions: []Faction{FctGER}, Category: WcHullMachinegun},
	WI_MG_42_Sd_Kfz_251_Half_track:       {ID: WI_MG_42_Sd_Kfz_251_Half_track, Name: "GER Half-track MG [GER Half-track]", Factions: []Faction{FctGER}, Category: WcMountedMachinegun},

	WI_MOSIN_NAGANT_1891:         {ID: WI_MOSIN_NAGANT_1891, Name: "Mosin-Nagant 1891", Factions: []Faction{FctRUS}, Category: WcBoltActionRifle},
	WI_MOSIN_NAGANT_91_30:        {ID: WI_MOSIN_NAGANT_91_30, Name: "Mosin-Nagant 91/30", Factions: []Faction{FctRUS}, Category: WcBoltActionRifle},
	WI_MOSIN_NAGANT_M38:          {ID: WI_MOSIN_NAGANT_M38, Name: "Mosin-Nagant M38", Factions: []Faction{FctRUS}, Category: WcBoltActionRifle},
	WI_SVT40:                     {ID: WI_SVT40, Name: "SVT40", Factions: []Faction{FctRUS}, Category: WcSemiAutoRifle},
	WI_PPSH_41:                   {ID: WI_PPSH_41, Name: "PPSh-41", Factions: []Faction{FctRUS}, Category: WcSubmachineGun},
	WI_PPSH_41_W_DRUM:            {ID: WI_PPSH_41_W_DRUM, Name: "PPSh-41 Drum", Factions: []Faction{FctRUS}, Category: WcSubmachineGun},
	WI_DP_27:                     {ID: WI_DP_27, Name: "DP-27", Factions: []Faction{FctRUS}, Category: WcMachinegun},
	WI_SCOPED_MOSIN_NAGANT_91_30: {ID: WI_SCOPED_MOSIN_NAGANT_91_30, Name: "Mosin-Nagant 91/30 (4x)", Factions: []Faction{FctRUS}, Category: WcSniperRifle},
	WI_SCOPED_SVT40:              {ID: WI_SCOPED_SVT40, Name: "SVT40 (4x)", Factions: []Faction{FctRUS}, Category: WcSniperRifle},
	WI_NAGANT_M1895:              {ID: WI_NAGANT_M1895, Name: "Nagant M1895", Factions: []Faction{FctRUS}, Category: WcPistol},
	WI_TOKAREV_TT33:              {ID: WI_TOKAREV_TT33, Name: "Tokarev TT33", Factions: []Faction{FctRUS}, Category: WcPistol},
	WI_MPL_50_SPADE:              {ID: WI_MPL_50_SPADE, Name: "RUS Melee", Factions: []Faction{FctRUS}, Category: WcMeele},
	WI_SATCHEL_CHARGE:            {ID: WI_SATCHEL_CHARGE, Name: "Satchel Charge", Factions: []Faction{FctRUS}, Category: WcSatchel},
	WI_RG_42_GRENADE:             {ID: WI_RG_42_GRENADE, Name: "RUS Grenade", Factions: []Faction{FctRUS}, Category: WcGrenade},
	WI_MOLOTOV:                   {ID: WI_MOLOTOV, Name: "Molotov", Factions: []Faction{FctRUS}, Category: WcGrenade},
	WI_PTRS_41:                   {ID: WI_PTRS_41, Name: "PTRS-41", Factions: []Faction{FctRUS}, Category: WcAntiTankRifle},
	WI_POMZ_AP_MINE:              {ID: WI_POMZ_AP_MINE, Name: "RUS AP Mine", Factions: []Faction{FctRUS}, Category: WcAntiPersonnelMine},
	WI_TM_35_AT_MINE:             {ID: WI_TM_35_AT_MINE, Name: "RUS AT Mine", Factions: []Faction{FctRUS}, Category: WcAntiTankMine},
	WI_57MM_CANNON_ZiS_2:         {ID: WI_57MM_CANNON_ZiS_2, Name: "RUS AT Gun", Factions: []Faction{FctRUS}, Category: WcAntiTankGun},
	WI_122MM_HOWITZER_M1938_M_30: {ID: WI_122MM_HOWITZER_M1938_M_30, Name: "RUS Artillery", Factions: []Faction{FctRUS}, Category: WcArtillerygun},
	WI_BA_10:                     {ID: WI_BA_10, Name: "RUS Roadkill [BA-10]", Factions: []Faction{FctRUS}, Category: WcVehicle},
	WI_T70:                       {ID: WI_T70, Name: "RUS Roadkill [T70]", Factions: []Faction{FctRUS}, Category: WcVehicle},
	WI_T34_76:                    {ID: WI_T34_76, Name: "RUS Roadkill [T34/76]", Factions: []Faction{FctRUS}, Category: WcVehicle},
	WI_IS_1:                      {ID: WI_IS_1, Name: "RUS Roadkill [IS-1]", Factions: []Faction{FctRUS}, Category: WcVehicle},
	WI_ZIS_5_Supply:              {ID: WI_ZIS_5_Supply, Name: "RUS Roadkill [RUS Supply Truck]", Factions: []Faction{FctRUS}, Category: WcVehicle},
	WI_ZIS_5_Transport:           {ID: WI_ZIS_5_Transport, Name: "RUS Roadkill [RUS Transport Truck]", Factions: []Faction{FctRUS}, Category: WcVehicle},
	WI_GAZ_67:                    {ID: WI_GAZ_67, Name: "RUS Roadkill [RUS Jeep]", Factions: []Faction{FctRUS}, Category: WcVehicle},
	WI_19_K_45MM_BA_10:           {ID: WI_19_K_45MM_BA_10, Name: "RUS Tank Cannon [BA-10]", Factions: []Faction{FctRUS}, Category: WcMainCannon},
	WI_COAXIAL_DT_BA_10:          {ID: WI_COAXIAL_DT_BA_10, Name: "RUS Tank Coaxial [BA-10]", Factions: []Faction{FctRUS}, Category: WcCoaxialMachinegun},
	WI_45MM_M1937_T70:            {ID: WI_45MM_M1937_T70, Name: "RUS Tank Cannon [T70]", Factions: []Faction{FctRUS}, Category: WcMainCannon},
	WI_COAXIAL_DT_T70:            {ID: WI_COAXIAL_DT_T70, Name: "RUS Tank Coaxial [T70]", Factions: []Faction{FctRUS}, Category: WcCoaxialMachinegun},
	WI_76MM_ZiS_5_T34_76:         {ID: WI_76MM_ZiS_5_T34_76, Name: "RUS Tank Cannon [T34/76]", Factions: []Faction{FctRUS}, Category: WcMainCannon},
	WI_COAXIAL_DT_T34_76:         {ID: WI_COAXIAL_DT_T34_76, Name: "RUS Tank Coaxial [T34/76]", Factions: []Faction{FctRUS}, Category: WcCoaxialMachinegun},
	WI_HULL_DT_T34_76:            {ID: WI_HULL_DT_T34_76, Name: "RUS Tank Hull MG [T34/76]", Factions: []Faction{FctRUS}, Category: WcHullMachinegun},
	WI_D_5T_85MM_IS_1:            {ID: WI_D_5T_85MM_IS_1, Name: "RUS Tank Cannon [IS-1]", Factions: []Faction{FctRUS}, Category: WcMainCannon},
	WI_COAXIAL_DT_IS_1:           {ID: WI_COAXIAL_DT_IS_1, Name: "RUS Tank Coaxial [IS-1]", Factions: []Faction{FctRUS}, Category: WcCoaxialMachinegun},
	WI_HULL_DT_IS_1:              {ID: WI_HULL_DT_IS_1, Name: "RUS Tank Hull MG [IS-1]", Factions: []Faction{FctRUS}, Category: WcHullMachinegun},

	WI_SMLE_No_1_Mk_III:                     {ID: WI_SMLE_No_1_Mk_III, Name: "SMLE Mk III", Factions: []Faction{FctGB}, Category: WcBoltActionRifle},
	WI_Rifle_No_5_Mk_I:                      {ID: WI_Rifle_No_5_Mk_I, Name: "Jungle Carbine", Factions: []Faction{FctGB}, Category: WcBoltActionRifle},
	WI_Rifle_No_4_Mk_I:                      {ID: WI_Rifle_No_4_Mk_I, Name: "No.4 Rifle Mk I", Factions: []Faction{FctGB}, Category: WcBoltActionRifle},
	WI_Sten_Gun_Mk_II:                       {ID: WI_Sten_Gun_Mk_II, Name: "Sten Gun Mk.II", Factions: []Faction{FctGB}, Category: WcSubmachineGun},
	WI_Sten_Gun_Mk_V:                        {ID: WI_Sten_Gun_Mk_V, Name: "Sten Gun Mk.V", Factions: []Faction{FctGB}, Category: WcSubmachineGun},
	WI_Lanchester:                           {ID: WI_Lanchester, Name: "Lanchester", Factions: []Faction{FctGB}, Category: WcSubmachineGun},
	WI_M1928A1_THOMPSON:                     {ID: WI_M1928A1_THOMPSON, Name: "M1928A1 Thompson", Factions: []Faction{FctGB}, Category: WcSubmachineGun},
	WI_Bren_Gun:                             {ID: WI_Bren_Gun, Name: "Bren Gun", Factions: []Faction{FctGB}, Category: WcAssaultRifle},
	WI_Lewis_Gun:                            {ID: WI_Lewis_Gun, Name: "Lewis Gun", Factions: []Faction{FctGB}, Category: WcMachinegun},
	WI_FLAMETHROWER:                         {ID: WI_FLAMETHROWER, Name: "GB Flamethrower", Factions: []Faction{FctGB}, Category: WcFlamethrower},
	WI_Lee_Enfield_Pattern_1914_Sniper:      {ID: WI_Lee_Enfield_Pattern_1914_Sniper, Name: "P14 Enfield (8x)", Factions: []Faction{FctGB}, Category: WcSniperRifle},
	WI_Rifle_No_4_Mk_I_Sniper:               {ID: WI_Rifle_No_4_Mk_I_Sniper, Name: "No.4 Mk I Sniper", Factions: []Faction{FctGB}, Category: WcSniperRifle},
	WI_Webley_MK_VI:                         {ID: WI_Webley_MK_VI, Name: "Webley Mk IV", Factions: []Faction{FctGB}, Category: WcPistol},
	WI_Fairbairn_Sykes:                      {ID: WI_Fairbairn_Sykes, Name: "GB Melee", Factions: []Faction{FctGB}, Category: WcMeele},
	WI_Satchel:                              {ID: WI_Satchel, Name: "Satchel Charge", Factions: []Faction{FctGB}, Category: WcSatchel},
	WI_Mills_Bomb:                           {ID: WI_Mills_Bomb, Name: "GB Grenade", Factions: []Faction{FctGB}, Category: WcGrenade},
	WI_NO_82_Grenade:                        {ID: WI_NO_82_Grenade, Name: "GB Grenade", Factions: []Faction{FctGB}, Category: WcGrenade},
	WI_PIAT:                                 {ID: WI_PIAT, Name: "PIAT", Factions: []Faction{FctGB}, Category: WcAntiTankRifle},
	WI_Boys_Anti_tank_Rifle:                 {ID: WI_Boys_Anti_tank_Rifle, Name: "Boys AT Rifle", Factions: []Faction{FctGB}, Category: WcAntiTankRifle},
	WI_A_P_Shrapnel_Mine_Mk_II:              {ID: WI_A_P_Shrapnel_Mine_Mk_II, Name: "GB AP Mine", Factions: []Faction{FctGB}, Category: WcAntiPersonnelMine},
	WI_A_T_Mine_G_S_Mk_V:                    {ID: WI_A_T_Mine_G_S_Mk_V, Name: "GB AT Mine", Factions: []Faction{FctGB}, Category: WcAntiTankMine},
	WI_QF_6_POUNDER_QF_6_Pounder:            {ID: WI_QF_6_POUNDER_QF_6_Pounder, Name: "GB AT Gun", Factions: []Faction{FctGB}, Category: WcAntiTankGun},
	WI_QF_25_POUNDER_QF_25_Pounder:          {ID: WI_QF_25_POUNDER_QF_25_Pounder, Name: "GB Artillery", Factions: []Faction{FctGB}, Category: WcArtillerygun},
	WI_Daimler:                              {ID: WI_Daimler, Name: "GB Roadkill [Daimler]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Tetrarch:                             {ID: WI_Tetrarch, Name: "GB Roadkill [Tetrarch]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_M3_Stuart_Honey:                      {ID: WI_M3_Stuart_Honey, Name: "GB Roadkill [Stuart Honey]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Cromwell:                             {ID: WI_Cromwell, Name: "GB Roadkill [Cromwell]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Crusader_Mk_III:                      {ID: WI_Crusader_Mk_III, Name: "GB Roadkill [Crusader]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Firefly:                              {ID: WI_Firefly, Name: "GB Roadkill [Firefly]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Churchill_Mk_III:                     {ID: WI_Churchill_Mk_III, Name: "GB Roadkill [Churchill]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Churchill_Mk_VII:                     {ID: WI_Churchill_Mk_VII, Name: "GB Roadkill [Churchill]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Bedford_OYD_Supply:                   {ID: WI_Bedford_OYD_Supply, Name: "GB Roadkill [GB Supply Truck]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_Bedford_OYD_Transport:                {ID: WI_Bedford_OYD_Transport, Name: "GB Roadkill [GB Transport Truck]", Factions: []Faction{FctGB}, Category: WcVehicle},
	WI_QF_2_POUNDER_Daimler:                 {ID: WI_QF_2_POUNDER_Daimler, Name: "GB Tank Cannon [Daimler]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_BESA_Daimler:                 {ID: WI_COAXIAL_BESA_Daimler, Name: "GB Tank Coaxial [Daimler]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_QF_2_POUNDER_Tetrarch:                {ID: WI_QF_2_POUNDER_Tetrarch, Name: "GB Tank Cannon [Tetrarch]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_BESA_Tetrarch:                {ID: WI_COAXIAL_BESA_Tetrarch, Name: "GB Tank Coaxial [Tetrarch]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_37MM_CANNON_M3_Stuart_Honey:          {ID: WI_37MM_CANNON_M3_Stuart_Honey, Name: "GB Tank Cannon [Stuart Honey]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_M1919_M3_Stuart_Honey:        {ID: WI_COAXIAL_M1919_M3_Stuart_Honey, Name: "GB Tank Coaxial [Stuart Honey]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_HULL_M1919_M3_Stuart_Honey:           {ID: WI_HULL_M1919_M3_Stuart_Honey, Name: "GB Tank Hull MG [Stuart Honey]", Factions: []Faction{FctGB}, Category: WcHullMachinegun},
	WI_OQF_75MM_Cromwell:                    {ID: WI_OQF_75MM_Cromwell, Name: "GB Tank Cannon [Cromwell]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_BESA_Cromwell:                {ID: WI_COAXIAL_BESA_Cromwell, Name: "GB Tank Coaxial [Cromwell]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_HULL_BESA_Cromwell:                   {ID: WI_HULL_BESA_Cromwell, Name: "GB Tank Hull MG [Cromwell]", Factions: []Faction{FctGB}, Category: WcHullMachinegun},
	WI_OQF_57MM_Crusader_Mk_III:             {ID: WI_OQF_57MM_Crusader_Mk_III, Name: "GB Tank Cannon [Crusader]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_BESA_Crusader_Mk_III:         {ID: WI_COAXIAL_BESA_Crusader_Mk_III, Name: "GB Tank Coaxial [Crusader]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_QF_17_POUNDER_Firefly:                {ID: WI_QF_17_POUNDER_Firefly, Name: "GB Tank Cannon [Firefly]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_M1919_Firefly:                {ID: WI_COAXIAL_M1919_Firefly, Name: "GB Tank Coaxial [Firefly]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_OQF_57MM_Churchill_Mk_III:            {ID: WI_OQF_57MM_Churchill_Mk_III, Name: "GB Tank Cannon [Churchill]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_BESA_7_92mm_Churchill_Mk_III: {ID: WI_COAXIAL_BESA_7_92mm_Churchill_Mk_III, Name: "GB Tank Coaxial [Churchill]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_HULL_BESA_7_92mm_Churchill_Mk_III:    {ID: WI_HULL_BESA_7_92mm_Churchill_Mk_III, Name: "GB Tank Hull MG [Churchill]", Factions: []Faction{FctGB}, Category: WcHullMachinegun},
	WI_OQF_75MM_Churchill_Mk_VII:            {ID: WI_OQF_75MM_Churchill_Mk_VII, Name: "GB Tank Cannon [Churchill]", Factions: []Faction{FctGB}, Category: WcMainCannon},
	WI_COAXIAL_BESA_7_92mm_Churchill_Mk_VII: {ID: WI_COAXIAL_BESA_7_92mm_Churchill_Mk_VII, Name: "GB Tank Coaxial [Churchill]", Factions: []Faction{FctGB}, Category: WcCoaxialMachinegun},
	WI_HULL_BESA_7_92mm_Churchill_Mk_VII:    {ID: WI_HULL_BESA_7_92mm_Churchill_Mk_VII, Name: "GB Tank Hull MG [Churchill]", Factions: []Faction{FctGB}, Category: WcHullMachinegun},

	WI_UNKNOWN:          {ID: WI_UNKNOWN, Name: "Unknown", Factions: []Faction{FctGER, FctUS, FctRUS, FctGB}, Category: WcUnknown},
	WI_BOMBING_RUN:      {ID: WI_BOMBING_RUN, Name: "Bombing Run", Factions: []Faction{FctGER, FctUS, FctGB}, Category: WcCommanderAbility},
	WI_STRAFING_RUN:     {ID: WI_STRAFING_RUN, Name: "Strafing Run", Factions: []Faction{FctGER, FctUS, FctRUS, FctGB}, Category: WcCommanderAbility},
	WI_PRECISION_STRIKE: {ID: WI_PRECISION_STRIKE, Name: "Precision Strike", Factions: []Faction{FctGER, FctUS, FctRUS, FctGB}, Category: WcCommanderAbility},
	WI_Unknown:          {ID: WI_Unknown, Name: "Katyusha Barrage", Factions: []Faction{FctRUS}, Category: WcCommanderAbility},
	WI_FLARE_GUN:        {ID: WI_FLARE_GUN, Name: "Flare Gun", Factions: []Faction{FctGER, FctUS, FctRUS, FctGB}, Category: WcFlaregun},
}

var fallback_weapon = Weapon{ID: WI_INVALID, Name: "Invalid", Factions: []Faction{FctGER, FctUS, FctRUS, FctGB}, Category: WcUnknown}

func ParseWeapon(weaponIdentifier string) Weapon {
	wi := WeaponIdentifier(weaponIdentifier)
	if weapon, ok := weapons[wi]; ok {
		return weapon
	}
	for _, v := range weapons {
		if strings.HasPrefix(string(v.ID), weaponIdentifier) {
			logger.Info("Using", v.ID, "as fallback for", weaponIdentifier)
			return v
		}
	}
	logger.Error("Weapon unparseable:", weaponIdentifier)
	return fallback_weapon
}
