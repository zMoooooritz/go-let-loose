---@meta

---Player information
---@class PlayerInfo
---@field Name string The player's display name
---@field ID string The player's unique identifier (Steam ID, etc.)
local PlayerInfo = {}

---Detailed player information with additional stats and status
---@class DetailedPlayerInfo : PlayerInfo
---@field Name string The player's display name
---@field ID string The player's unique identifier
---@field ClanTag string The player's clan tag
---@field Platform string Player's platform ("steam", "epic", "xbl", "none")
---@field EOSID string Player's Epic Online Services ID
---@field Team number The player's team (0=Allies, 1=Axis, 2=None)
---@field Role number The player's role ID
---@field Platoon string The player's platoon/squad
---@field Stats PlayerStats Player's stats
---@field ScoreData Score Player's score breakdown
---@field Loadout string The player's current loadout
---@field Level number Player's level
---@field WorldPosition Position Player's world position
local DetailedPlayerInfo = {}

---Weapon information
---@class Weapon
---@field ID string Weapon identifier (e.g., "M1_GARAND")
---@field Name string Human-readable weapon name (e.g., "M1 Garand")
---@field Category string Weapon category (Rifle, SMG, etc.)
---@field Factions table Array of factions that can use this weapon
local Weapon = {}

---Unit/Squad information
---@class Unit
---@field Name string Unit name (e.g., "Able", "Baker", "Command")
---@field ID number Unit ID
local Unit = {}

---Player statistics
---@class PlayerStats
---@field Deaths number Player's death count
---@field InfantryKills number Infantry kill count
---@field VehicleKills number Vehicle kill count
---@field TeamKills number Team kill count
---@field VehiclesDestroyed number Vehicles destroyed count
local PlayerStats = {}

---Player score breakdown
---@class Score
---@field Combat number Combat score
---@field Offense number Offensive score
---@field Defense number Defensive score
---@field Support number Support score
local Score = {}

---3D position in game world
---@class Position
---@field X number X coordinate
---@field Y number Y coordinate
---@field Z number Z coordinate (elevation)
local Position = {}

---Team score data
---@class TeamData
---@field Allies number Allied team score/count
---@field Axis number Axis team score/count
local TeamData = {}

---Game map information
---@class GameMap
---@field ID string Map identifier
---@field Name string Full map name
---@field Tag string Short map tag
---@field PrettyName string Display name
---@field ShortName string Abbreviated name
---@field Allies string Allied faction
---@field Axis string Axis faction
---@field Orientation string Map orientation
---@field MirroredFactions boolean Whether factions are mirrored on this map
local GameMap = {}

---Layer information (map + game mode)
---@class Layer
---@field ID string Layer identifier
---@field GameMap GameMap Map information
---@field GameMode string Game mode
---@field Attackers string Attacking team
---@field Environment string Environmental conditions
local Layer = {}

---Current game state
---@class GameState
---@field PlayerCount TeamData Player count per team
---@field GameScore TeamData Current scores
---@field RemainingSeconds number Time remaining in match
---@field CurrentMap Layer Current map/layer
---@field NextMap Layer Next map/layer
local GameState = {}

---Admin information
---@class Admin
---@field UserId string Admin's player ID
---@field Group string Admin group
---@field Comment string Admin comment/description
local Admin = {}

---VIP player information
---@class VipPlayerEntry
---@field PlayerId string VIP player's ID
---@field Comment string VIP comment/description
local VipPlayerEntry = {}


---Server ban information
---@class ServerBan
---@field UserId string Banned player's user ID
---@field UserName string Banned player's name
---@field TimeOfBanning string Ban timestamp
---@field DurationHours number Ban duration in hours (for temp bans)
---@field BanReason string Ban reason
---@field AdminName string Admin who issued the ban
local ServerBan = {}

---Server configuration
---@class ServerConfig
---@field Name string Server name
---@field BuildNumber string Build number
---@field BuildRevision string Build revision
---@field SupportedPlatforms table Array of supported platforms
---@field PasswordProtected boolean Whether server is password protected
local ServerConfig = {}

---Session information
---@class SessionInfo
---@field ServerName string Server name
---@field MapName string Current map name
---@field MapId string Current map ID
---@field GameMode string Current game mode
---@field RemainingMatchTime number Time remaining in match (seconds)
---@field MatchTime number Total match time (seconds)
---@field AlliedFaction number Allied faction ID
---@field AxisFaction number Axis faction ID
---@field MaxPlayerCount number Maximum players
---@field AlliedScore number Allied team score
---@field AxisScore number Axis team score
---@field PlayerCount number Current player count
---@field AlliedPlayerCount number Allied player count
---@field AxisPlayerCount number Axis player count
---@field MaxQueueCount number Maximum queue size
---@field QueueCount number Current queue size
---@field MaxVipQueueCount number Maximum VIP queue size
---@field VipQueueCount number Current VIP queue size
local SessionInfo = {}

---Command information
---@class Command
---@field ID string Command identifier
---@field Name string Command name
---@field ClientSupported boolean Whether command is client-supported
local Command = {}

---Command details with parameters
---@class CommandDetails
---@field Name string Command name
---@field Text string Command text
---@field Description string Command description
---@field DialogueParameters table Array of DialogueParameter
local CommandDetails = {}

---Map information in rotation/sequence
---@class MapInformation
---@field Name string Map name
---@field GameMode string Game mode
---@field TimeOfDay string Time of day setting
---@field Id string Map identifier
---@field Position number Position in rotation
local MapInformation = {}

---Map rotation response
---@class MapRotation
---@field CurrentIndex number Current map index in rotation
---@field Maps MapInformation[] Array of maps in rotation
local MapRotation = {}

---Map sequence response
---@class MapSequence
---@field Maps MapInformation[] Array of maps in sequence
local MapSequence = {}

---Dialogue parameter for commands
---@class DialogueParameter
---@field Type string Parameter type
---@field Name string Parameter name
---@field ID string Parameter ID
---@field DisplayMember string Display member
---@field ValueMember string Value member
local DialogueParameter = {}

---Server view with team and squad information
---@class ServerView
---@field Allies TeamView Allied team view
---@field Axis TeamView Axis team view
---@field Neutral SquadView Neutral players
local ServerView = {}

---Team view with commander and squads
---@class TeamView
---@field Commander DetailedPlayerInfo Team commander
---@field Squads table Map of squad name to SquadView
local TeamView = {}

---Squad view with players
---@class SquadView
---@field Team string Squad's team
---@field SquadType string Squad type ("Infantry", "Recon", "Armor")
---@field Name string Squad name
---@field Players table Array of DetailedPlayerInfo
local SquadView = {}

---Log entry
---@class LogEntry
---@field Timestamp string Log timestamp
---@field Level string Log level
---@field Message string Log message
---@field Raw string Raw log line
local LogEntry = {}

---Admin log entry
---@class AdminLogEntry
---@field Timestamp string Log timestamp (RFC3339 format)
---@field Message string Log message
local AdminLogEntry = {}
