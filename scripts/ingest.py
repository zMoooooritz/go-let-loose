#!/usr/bin/env python3

"""
Go code generator for Hell Let Loose game data.
Converts JSON data into Go map initializations.
"""

import json
import re
from typing import Any, Dict, List

# ============================================================================
# CONSTANTS - Mapping game data IDs to Go constants
# ============================================================================

FACTION_MAP = {
    0: "FACTION_GER",
    1: "FACTION_US",
    2: "FACTION_SOV",
    3: "FACTION_CW",
    4: "FACTION_DAK",
    5: "FACTION_B8A"
}

FACTION_SHORT_MAP = {
    0: "GER",
    1: "US",
    2: "SOV",
    3: "CW",
    4: "DAK",
    5: "B8A"
}

ROLE_MAP = {
    0: "Rifleman",
    1: "Assault",
    2: "AutomaticRifleman",
    3: "Medic",
    4: "Spotter",
    5: "Support",
    6: "HeavyMachinegunner",
    7: "AntiTank",
    8: "Engineer",
    9: "Officer",
    10: "Sniper",
    11: "Crewman",
    12: "TankCommander",
    13: "ArmyCommander",
    14: "ArtilleryObserver",
    15: "Operator",
    16: "Gunner"
}

TEAM_MAP = {
    1: "TEAM_ALLIES",
    2: "TEAM_AXIS"
}

# ============================================================================
# UTILITY FUNCTIONS
# ============================================================================

def to_identifier(prefix: str, name: str) -> str:
    """Convert a name to a Go constant identifier with prefix."""
    # Special cases
    if prefix == "WEAPON" and name == "Unknown":
        identifier = "Unknown"
    elif prefix == "WEAPON" and name == "Satchel":
        identifier = "Satchel"
    elif prefix == "WEAPON" and name == "QF 25 POUNDER":
        identifier = "QF_25_POUNDER_GUN"
    else:
        identifier = re.sub(r'[^A-Z0-9]', '_', name.upper())
        identifier = re.sub(r'__+', '_', identifier)
        identifier = identifier.strip('_')
    
    return f"{prefix}_{identifier}"

def indent(level: int) -> str:
    """Return indentation string for given level."""
    return "    " * level

# ============================================================================
# FIELD RENDERING - Generic rendering of Go struct fields
# ============================================================================

class FieldRenderer:
    """Handles rendering of different field types in Go structs."""
    
    def render_string(self, name: str, value: str, depth: int = 2) -> str:
        """Render a string field."""
        return f'{indent(depth)}{name}: "{value}",'
    
    def render_int(self, name: str, value: int, depth: int = 2) -> str:
        """Render an integer field."""
        return f'{indent(depth)}{name}: {value},'
    
    def render_bool(self, name: str, value: bool, depth: int = 2) -> str:
        """Render a boolean field."""
        go_bool = "true" if value else "false"
        return f'{indent(depth)}{name}: {go_bool},'
    
    def render_const(self, name: str, value: str, depth: int = 2) -> str:
        """Render a constant (no quotes)."""
        return f'{indent(depth)}{name}: {value},'
    
    def render_faction_list(self, name: str, factions: List[Dict], depth: int = 2) -> str:
        """Render a list of faction constants."""
        faction_consts = [FACTION_MAP[f['id']] for f in factions]
        return f'{indent(depth)}{name}: []Faction{{{", ".join(faction_consts)}}},'
    
    def render_role_list(self, name: str, roles: List[Dict], depth: int = 2) -> str:
        """Render a list of role constants."""
        if not roles:
            return f'{indent(depth)}{name}: []RoleIdentifier{{}},'
        role_consts = [to_identifier('ROLE', ROLE_MAP[r['id']]) for r in roles]
        return f'{indent(depth)}{name}: []RoleIdentifier{{{", ".join(role_consts)}}},'
    
    def render_weapon_id_list(self, name: str, weapons: List[str], depth: int = 2) -> str:
        """Render a list of weapon identifier constants."""
        if not weapons:
            return f'{indent(depth)}{name}: []WeaponIdentifier{{}},'
        weapon_consts = [to_identifier('WEAPON', w) for w in weapons]
        return f'{indent(depth)}{name}: []WeaponIdentifier{{{", ".join(weapon_consts)}}},'
    
    def render_loadout_items(self, name: str, items: List[Dict], depth: int = 2) -> str:
        """Render a list of loadout items."""
        if not items:
            return f'{indent(depth)}{name}: []LoadoutItem{{}},'
        
        lines = [f'{indent(depth)}{name}: []LoadoutItem{{']
        for item in items:
            lines.append(f'{indent(depth+1)}{{Name: "{item["name"]}", Amount: {item["amount"]}}},')
        lines.append(f'{indent(depth)}}},')
        return '\n'.join(lines)

renderer = FieldRenderer()

# ============================================================================
# CONST GENERATION - Generate Go const declarations
# ============================================================================

def print_const_block(const_list: List[tuple], type_name: str):
    """
    Print a Go const block.
    
    Args:
        const_list: List of (const_name, const_value) tuples
        type_name: Go type for the constants
    """
    if not const_list:
        return
    
    print("const (")
    for const_name, const_value in sorted(const_list, key=lambda x: x[0]):
        print(f'    {const_name} {type_name} = "{const_value}"')
    print(")")
    print()

def extract_unique_values(data: Dict, key: str) -> set:
    """Extract unique values for a given key from all data items."""
    values = set()
    for item_data in data.values():
        if key in item_data:
            values.add(item_data[key])
    return values

# ============================================================================
# GENERIC MAP GENERATION
# ============================================================================

def print_go_map_header(map_name: str, key_type: str, value_type: str):
    """Print the header of a Go map."""
    print(f"var {map_name} = map[{key_type}]{value_type}{{")

def print_go_map_footer():
    """Print the footer of a Go map."""
    print("}")

def print_struct_field(field_name: str, field_value: Any, field_type: str, depth: int = 2):
    """Print a single struct field based on its type."""
    if field_type == 'string':
        print(renderer.render_string(field_name, field_value, depth))
    elif field_type == 'int':
        print(renderer.render_int(field_name, field_value, depth))
    elif field_type == 'bool':
        print(renderer.render_bool(field_name, field_value, depth))
    elif field_type == 'const':
        print(renderer.render_const(field_name, field_value, depth))
    elif field_type == 'faction_list':
        print(renderer.render_faction_list(field_name, field_value, depth))
    elif field_type == 'role_list':
        print(renderer.render_role_list(field_name, field_value, depth))
    elif field_type == 'weapon_id_list':
        print(renderer.render_weapon_id_list(field_name, field_value, depth))
    elif field_type == 'loadout_items':
        print(renderer.render_loadout_items(field_name, field_value, depth))
    else:
        # Default: try to print as-is
        print(f'{indent(depth)}{field_name}: {field_value},')

# ============================================================================
# ENTITY-SPECIFIC GENERATORS
# ============================================================================

def render_vehicle_seat(seat: Dict, depth: int = 3):
    """Render a single vehicle seat struct."""
    print(f'{indent(depth)}{{')
    print_struct_field('Index', seat['index'], 'int', depth + 1)
    
    seat_type_const = to_identifier("VEHICLE_SEAT_TYPE", seat['type'])
    print_struct_field('Type', seat_type_const, 'const', depth + 1)
    
    # Weapons
    weapon_ids = [w['id'] for w in seat['weapons']] if seat['weapons'] else []
    print_struct_field('Weapons', weapon_ids, 'weapon_id_list', depth + 1)
    
    # Required roles
    print_struct_field('RequiresRoles', seat.get('requires_roles', []), 'role_list', depth + 1)
    
    # Exposed
    print_struct_field('Exposed', seat['exposed'], 'bool', depth + 1)
    
    print(f'{indent(depth)}}},')

def do_vehicle(vehicles: Dict):
    """Generate Go code for vehicle identifiers, types, and map."""
    # Generate vehicle identifier constants
    vehicle_consts = [(to_identifier("VEHICLE", vid), vid) for vid in vehicles.keys()]
    print_const_block(vehicle_consts, "VehicleIdentifier")
    
    # Generate vehicle type constants
    vehicle_types = extract_unique_values(vehicles, 'type')
    type_consts = [(to_identifier("VEHICLE_TYPE", vt), vt) for vt in vehicle_types]
    print_const_block(type_consts, "VehicleType")
    
    # Generate vehicle seat type constants
    seat_types = set()
    for v in vehicles.values():
        for seat in v.get('seats', []):
            seat_types.add(seat['type'])
    seat_type_consts = [(to_identifier("VEHICLE_SEAT_TYPE", st), st) for st in seat_types]
    print_const_block(seat_type_consts, "VehicleSeatType")
    
    # Generate the map
    print_go_map_header("vehicleMap", "VehicleIdentifier", "Vehicle")
    
    for vehicle_id, vehicle_data in vehicles.items():
        const_name = to_identifier("VEHICLE", vehicle_id)
        print(f"    {const_name}: {{")
        
        # Basic fields
        print_struct_field('ID', const_name, 'const')
        print_struct_field('Name', vehicle_data['name'], 'string')
        print_struct_field('Factions', vehicle_data['factions'], 'faction_list')
        
        # Type
        type_const = to_identifier("VEHICLE_TYPE", vehicle_data['type'])
        print_struct_field('Type', type_const, 'const')
        
        # Seats (complex nested structure)
        if vehicle_data['seats']:
            print(f"        Seats: []VehicleSeat{{")
            for seat in vehicle_data['seats']:
                render_vehicle_seat(seat)
            print(f"        }},")
        else:
            print(f"        Seats: []VehicleSeat{{}},")
        
        print(f"    }},")
    
    print_go_map_footer()

def do_weapon(weapons: Dict):
    """Generate Go code for weapon identifiers, types, and map."""
    # Generate weapon identifier constants
    weapon_consts = [(to_identifier("WEAPON", wid), wid) for wid in weapons.keys()]
    print_const_block(weapon_consts, "WeaponIdentifier")
    
    # Generate weapon type constants
    weapon_types = extract_unique_values(weapons, 'type')
    type_consts = [(to_identifier("WEAPON_TYPE", wt), wt) for wt in weapon_types]
    print_const_block(type_consts, "WeaponType")
    
    # Generate the map
    print_go_map_header("weaponMap", "WeaponIdentifier", "Weapon")
    
    for weapon_id, weapon_data in weapons.items():
        const_name = to_identifier("WEAPON", weapon_id)
        print(f"    {const_name}: {{")
        
        # Basic fields
        print_struct_field('ID', const_name, 'const')
        print_struct_field('Name', weapon_data['name'], 'string')
        
        # Type
        type_const = to_identifier("WEAPON_TYPE", weapon_data['type'])
        print_struct_field('Type', type_const, 'const')
        
        # Factions
        print_struct_field('Factions', weapon_data['factions'], 'faction_list')
        
        # Magnification (optional field)
        magnification = weapon_data.get('magnification', 0)
        print_struct_field('Magnification', magnification, 'int')
        
        print(f"    }},")
    
    print_go_map_footer()

def do_role(roles: Dict):
    """Generate Go code for role identifiers and map."""
    # Generate role identifier constants
    role_consts = [(to_identifier("ROLE", r['name']), r['name']) for r in roles.values()]
    print_const_block(role_consts, "RoleIdentifier")
    
    # Generate the map
    print_go_map_header("roleMap", "RoleIdentifier", "Role")
    
    for role_id, role_data in roles.items():
        const_name = to_identifier("ROLE", role_data['name'])
        print(f"    {const_name}: {{")
        
        # Basic fields
        print_struct_field('ID', role_data['id'], 'int')
        print_struct_field('Name', const_name, 'const')
        print_struct_field('PrettyName', role_data['pretty_name'], 'string')
        
        # Role type (Squad type)
        type_const = to_identifier("SQUAD_TYPE", role_data['type'])
        print_struct_field('RoleType', type_const, 'const')
        
        # Squad leader flag
        print_struct_field('IsSquadLeader', role_data['is_squad_leader'], 'bool')
        
        # Combat scores
        print_struct_field('KillCombatScore', role_data['kill_combat_score'], 'int')
        print_struct_field('AssistCombatScore', role_data['assist_combat_score'], 'int')
        
        print(f"    }},")
    
    print_go_map_footer()

def do_layer(layers: Dict):
    """Generate Go code for layer map."""
    # Generate layer identifier constants
    layer_consts = [(to_identifier("LAYER", l['id']), l['id']) for l in layers.values()]
    print_const_block(layer_consts, "LayerIdentifier")

    # Generate the map
    print_go_map_header("layerMap", "LayerIdentifier", "Layer")
    
    for layer_id, layer_data in layers.items():
        const_name = to_identifier("LAYER", layer_id)
        print(f"    {const_name}: {{")
        
        print_struct_field('ID', const_name, 'const')
        
        map_const = to_identifier("MAP", layer_data['map']['id'])
        print_struct_field('MapIdentifier', map_const, 'const')
        
        gamemode_const = to_identifier("GAMEMODE", layer_data['game_mode']['id'])
        print_struct_field('GameModeIdentifier', gamemode_const, 'const')
        
        tod_const = to_identifier("TOD", layer_data['time_of_day'])
        print_struct_field('TimeOfDay', tod_const, 'const')
        
        weather_const = to_identifier("WEATHER", layer_data['weather'])
        print_struct_field('Weather', weather_const, 'const')
        
        print_struct_field('PrettyName', layer_data['pretty_name'], 'string')
        
        attacking_team = TEAM_MAP.get(layer_data['attacking_team']['id'], 'TEAM_NONE') if layer_data['attacking_team'] else 'TEAM_NONE'
        print_struct_field('AttackingTeam', attacking_team, 'const')
        
        defending_team = TEAM_MAP.get(layer_data['defending_team']['id'], 'TEAM_NONE') if layer_data['defending_team'] else 'TEAM_NONE'
        print_struct_field('DefendingTeam', defending_team, 'const')
        
        attacking_faction = FACTION_MAP.get(layer_data['attacking_faction']['id'], 'FACTION_UNASSIGNED') if layer_data['attacking_faction'] else 'FACTION_UNASSIGNED'
        print_struct_field('AttackingFaction', attacking_faction, 'const')
        
        defending_faction = FACTION_MAP.get(layer_data['defending_faction']['id'], 'FACTION_UNASSIGNED') if layer_data['defending_faction'] else 'FACTION_UNASSIGNED'
        print_struct_field('DefendingFaction', defending_faction, 'const')
        
        print(f"    }},")
    
    print_go_map_footer()

def do_loadouts(loadouts: Dict):
    """Generate Go code for loadout constants and map."""
    # First, generate the constants
    loadout_ids = []
    for k, v in loadouts.items():
        faction_short = FACTION_SHORT_MAP[v['faction']['id']]
        role_name = v['role']['name']
        loadout_name = v['name']
        loadout_id = f"{faction_short}_{role_name}_{loadout_name}"
        loadout_ids.append(loadout_id)
    
    print("const (")
    for loadout_id in sorted(loadout_ids):
        print(f"    {to_identifier('LOADOUT', loadout_id)} LoadoutIdentifier = \"{loadout_id}\"")
    print(")")
    
    print()
    
    # Now generate the map
    print_go_map_header("loadoutMap", "LoadoutIdentifier", "Loadout")
    
    for loadout_key, loadout_data in loadouts.items():
        faction_id = loadout_data['faction']['id']
        role_name = loadout_data['role']['name']
        loadout_name = loadout_data['name']
        
        faction_short = FACTION_SHORT_MAP[faction_id]
        loadout_id = f"{faction_short}_{role_name}_{loadout_name}"
        const_name = to_identifier("LOADOUT", loadout_id)
        
        print(f"    {const_name}: {{")
        
        # Basic fields
        print_struct_field('ID', const_name, 'const')
        print_struct_field('Name', loadout_name, 'string')
        
        # Faction
        faction_const = FACTION_MAP[faction_id]
        print_struct_field('Faction', faction_const, 'const')
        
        # Role
        role_const = to_identifier("ROLE", role_name)
        print_struct_field('Role', role_const, 'const')
        
        # Required level
        print_struct_field('RequiredLevel', loadout_data['requires_level'], 'int')
        
        # Items (complex list)
        print_struct_field('Items', loadout_data['items'], 'loadout_items')
        
        print(f"    }},")
    
    print_go_map_footer()


# ============================================================================
# MAIN ENTRY POINT
# ============================================================================

def main():
    """Main entry point - load data and generate Go code."""
    with open('formatted_data.json', 'r') as f:
        data = json.load(f)
    
    # Uncomment the generators you want to use:
    # do_weapon(data.get('weapon', {}))
    # do_vehicle(data.get('vehicle', {}))
    # do_role(data.get('role', {}))
    # do_layer(data.get('layer', {}))
    # do_loadouts(data.get('loadout', {}))
    

if __name__ == "__main__":
    main()