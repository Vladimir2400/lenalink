#!/usr/bin/env python3
"""
Route Generator for LenaLink
Generates comprehensive routes for a full week (Nov 28 - Dec 4, 2025)
Includes bidirectional routes and multi-segment journeys
"""

from datetime import datetime, timedelta
import random

# Date range
START_DATE = datetime(2025, 11, 28)
DAYS = 7

# Base prices (will be varied by ±10% per day)
BASE_PRICES = {
    'moscow_yakutsk': 32000,
    'yakutsk_mirny': 12000,
    'yakutsk_neryungri': 14000,
    'yakutsk_udachny': 16000,
    'yakutsk_vilyuysk': 11000,
    'yakutsk_tiksi': 22000,
    'yakutsk_pokrovsk': 1200,
}

def format_date(dt):
    """Format date as MMDD for route IDs"""
    return dt.strftime('%m%d')

def format_datetime(dt):
    """Format datetime for SQL"""
    return dt.strftime('%Y-%m-%d %H:%M:%S')

def vary_price(base_price, variation=0.1):
    """Add random variation to price"""
    factor = 1 + random.uniform(-variation, variation)
    return round(base_price * factor, 2)

def generate_route_sql(route_id, from_city, to_city, departure, arrival, duration_ns,
                       price, reliability, insurance, transport_types):
    """Generate route INSERT statement"""
    return f"""INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('{route_id}', '{from_city}', '{to_city}', '{format_datetime(departure)}', '{format_datetime(arrival)}',
 {duration_ns}, {price:.2f}, {reliability:.2f}, {insurance:.2f}, false, ARRAY{transport_types}::text[], NOW());\n"""

def generate_segment_sql(seg_id, route_id, transport_type, provider, start_stop, end_stop,
                        departure, arrival, price, duration_ns, seats, reliability, distance, seq):
    """Generate segment INSERT statement"""
    return f"""INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('{seg_id}', '{route_id}', '{transport_type}', '{provider}', '{start_stop}', '{end_stop}',
 '{format_datetime(departure)}', '{format_datetime(arrival)}',
 {price:.2f}, {duration_ns}, {seats}, {reliability:.2f}, {distance}, {seq});\n"""

def generate_connection_sql(route_id, from_seg, to_seg, transfer_ns, distance, requires_transport, gap_ns, seq):
    """Generate connection INSERT statement"""
    return f"""INSERT INTO connections (route_id, from_segment_id, to_segment_id, transfer_duration,
                        transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('{route_id}', '{from_seg}', '{to_seg}',
 {transfer_ns}, {distance}, {str(requires_transport).lower()}, true, {gap_ns}, {seq});\n"""

def ns_to_hours(hours):
    """Convert hours to nanoseconds"""
    return int(hours * 3600 * 1e9)

def generate_day_routes(date):
    """Generate all routes for a single day"""
    date_str = format_date(date)
    routes = []
    segments = []
    connections = []

    # ========== Moscow ↔ Yakutsk (3 flights each direction) ==========
    msk_yks_flights = [
        ('morning', 8, 23, 'S7 Airlines', 'moscow_dme', 180, 92.0),
        ('afternoon', 14, 29, 'Aeroflot', 'moscow_svo', 220, 91.0),  # next day arrival
        ('evening', 18, 33, 'Yakutia Airlines', 'moscow_svo', 200, 90.0),  # next day arrival
    ]

    for time_slot, dep_hour, arr_hour, provider, airport, seats, reliability in msk_yks_flights:
        route_id = f'msk_yks_{date_str}_{time_slot}'
        seg_id = f'seg_msk_yks_{date_str}_{time_slot[0]}'

        departure = date.replace(hour=dep_hour, minute=0)
        # Handle next-day arrivals
        if arr_hour >= 24:
            arrival = (date + timedelta(days=1)).replace(hour=arr_hour-24, minute=0)
        else:
            arrival = date.replace(hour=arr_hour, minute=0)

        price = vary_price(BASE_PRICES['moscow_yakutsk'])
        insurance = price * 0.05

        routes.append(generate_route_sql(
            route_id, 'moscow', 'yakutsk', departure, arrival, ns_to_hours(8),
            price, reliability, insurance, "['air']"
        ))

        segments.append(generate_segment_sql(
            seg_id, route_id, 'air', provider, airport, 'yakutsk_yks',
            departure, arrival, price, ns_to_hours(8), seats, reliability, 4884, 1
        ))

    yks_msk_flights = [
        ('early', 1, 9, 'Yakutia Airlines', 'moscow_dme', 200, 91.0),
        ('morning', 6, 14, 'S7 Airlines', 'moscow_dme', 180, 91.0),
        ('evening', 19, 27, 'Aeroflot', 'moscow_svo', 220, 90.0),  # next day arrival
    ]

    for time_slot, dep_hour, arr_hour, provider, airport, seats, reliability in yks_msk_flights:
        route_id = f'yks_msk_{date_str}_{time_slot}'
        seg_id = f'seg_yks_msk_{date_str}_{time_slot[0:2]}'

        departure = date.replace(hour=dep_hour, minute=0)
        if arr_hour >= 24:
            arrival = (date + timedelta(days=1)).replace(hour=arr_hour-24, minute=0)
        else:
            arrival = date.replace(hour=arr_hour, minute=0)

        price = vary_price(BASE_PRICES['moscow_yakutsk'] * 1.08)  # Return flights slightly more expensive
        insurance = price * 0.05

        routes.append(generate_route_sql(
            route_id, 'yakutsk', 'moscow', departure, arrival, ns_to_hours(8),
            price, reliability, insurance, "['air']"
        ))

        segments.append(generate_segment_sql(
            seg_id, route_id, 'air', provider, 'yakutsk_yks', airport,
            departure, arrival, price, ns_to_hours(8), seats, reliability, 4884, 1
        ))

    # ========== Yakutsk ↔ Mirny (3 flights each direction) ==========
    mirny_flights = [
        ('morning', 9, 10.5),
        ('afternoon', 13, 14.5),
        ('evening', 17, 18.5),
    ]

    for time_slot, dep_hour, arr_hour in mirny_flights:
        # Yakutsk → Mirny
        route_id = f'yks_mirny_{date_str}_{time_slot}'
        seg_id = f'seg_yks_mirny_{date_str}_{time_slot[0]}'

        dep_minutes = int((dep_hour % 1) * 60)
        arr_minutes = int((arr_hour % 1) * 60)
        departure = date.replace(hour=int(dep_hour), minute=dep_minutes)
        arrival = date.replace(hour=int(arr_hour), minute=arr_minutes)

        price = vary_price(BASE_PRICES['yakutsk_mirny'])
        insurance = price * 0.05

        routes.append(generate_route_sql(
            route_id, 'yakutsk', 'mirny', departure, arrival, ns_to_hours(1.5),
            price, 88.0 + (time_slot == 'evening') * 2, insurance, "['air']"
        ))

        segments.append(generate_segment_sql(
            seg_id, route_id, 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
            departure, arrival, price, ns_to_hours(1.5), 100, 88.0, 520, 1
        ))

        # Mirny → Yakutsk (2 hours after arrival)
        route_id_back = f'mirny_yks_{date_str}_{time_slot}'
        seg_id_back = f'seg_mirny_yks_{date_str}_{time_slot[0]}'

        departure_back = arrival + timedelta(hours=0.5)
        arrival_back = departure_back + timedelta(hours=1.5)

        price_back = vary_price(BASE_PRICES['yakutsk_mirny'] * 1.05)
        insurance_back = price_back * 0.05

        routes.append(generate_route_sql(
            route_id_back, 'mirny', 'yakutsk', departure_back, arrival_back, ns_to_hours(1.5),
            price_back, 88.0, insurance_back, "['air']"
        ))

        segments.append(generate_segment_sql(
            seg_id_back, route_id_back, 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
            departure_back, arrival_back, price_back, ns_to_hours(1.5), 100, 88.0, 520, 1
        ))

    # ========== Yakutsk ↔ Neryungri (2 flights each direction) ==========
    ner_flights = [('morning', 10), ('evening', 16)]

    for time_slot, dep_hour in ner_flights:
        # Yakutsk → Neryungri
        route_id = f'yks_ner_{date_str}_{time_slot}'
        seg_id = f'seg_yks_ner_{date_str}_{time_slot[0]}'

        departure = date.replace(hour=dep_hour, minute=0)
        arrival = departure + timedelta(hours=1.5)

        price = vary_price(BASE_PRICES['yakutsk_neryungri'])
        insurance = price * 0.05

        routes.append(generate_route_sql(
            route_id, 'yakutsk', 'neryungri', departure, arrival, ns_to_hours(1.5),
            price, 87.0, insurance, "['air']"
        ))

        segments.append(generate_segment_sql(
            seg_id, route_id, 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
            departure, arrival, price, ns_to_hours(1.5), 120, 87.0, 560, 1
        ))

        # Neryungri → Yakutsk (2 hours after)
        route_id_back = f'ner_yks_{date_str}_{time_slot}'
        seg_id_back = f'seg_ner_yks_{date_str}_{time_slot[0]}'

        departure_back = arrival + timedelta(hours=0.5)
        arrival_back = departure_back + timedelta(hours=1.5)

        price_back = vary_price(BASE_PRICES['yakutsk_neryungri'] * 1.05)
        insurance_back = price_back * 0.05

        routes.append(generate_route_sql(
            route_id_back, 'neryungri', 'yakutsk', departure_back, arrival_back, ns_to_hours(1.5),
            price_back, 87.0, insurance_back, "['air']"
        ))

        segments.append(generate_segment_sql(
            seg_id_back, route_id_back, 'air', 'Yakutia Airlines', 'neryungri_aprt', 'yakutsk_yks',
            departure_back, arrival_back, price_back, ns_to_hours(1.5), 120, 87.0, 560, 1
        ))

    # ========== Yakutsk ↔ Udachny (1 flight each direction) ==========
    route_id = f'yks_udachny_{date_str}'
    seg_id = f'seg_yks_udachny_{date_str}'

    departure = date.replace(hour=11, minute=30)
    arrival = date.replace(hour=13, minute=45)
    price = vary_price(BASE_PRICES['yakutsk_udachny'])
    insurance = price * 0.05

    routes.append(generate_route_sql(
        route_id, 'yakutsk', 'udachny', departure, arrival, ns_to_hours(2.25),
        price, 85.0, insurance, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id, route_id, 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
        departure, arrival, price, ns_to_hours(2.25), 80, 85.0, 630, 1
    ))

    # Reverse
    route_id_back = f'udachny_yks_{date_str}'
    seg_id_back = f'seg_udachny_yks_{date_str}'

    departure_back = date.replace(hour=14, minute=30)
    arrival_back = date.replace(hour=16, minute=45)
    price_back = vary_price(BASE_PRICES['yakutsk_udachny'] * 1.03)
    insurance_back = price_back * 0.05

    routes.append(generate_route_sql(
        route_id_back, 'udachny', 'yakutsk', departure_back, arrival_back, ns_to_hours(2.25),
        price_back, 85.0, insurance_back, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id_back, route_id_back, 'air', 'ALROSA Air', 'udachny_aprt', 'yakutsk_yks',
        departure_back, arrival_back, price_back, ns_to_hours(2.25), 80, 85.0, 630, 1
    ))

    # ========== Yakutsk ↔ Vilyuysk (1 flight each direction) ==========
    route_id = f'yks_vilyuysk_{date_str}'
    seg_id = f'seg_yks_vilyuysk_{date_str}'

    departure = date.replace(hour=14, minute=0)
    arrival = date.replace(hour=15, minute=40)
    price = vary_price(BASE_PRICES['yakutsk_vilyuysk'])
    insurance = price * 0.05

    routes.append(generate_route_sql(
        route_id, 'yakutsk', 'vilyuysk', departure, arrival, ns_to_hours(1.67),
        price, 82.0, insurance, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id, route_id, 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
        departure, arrival, price, ns_to_hours(1.67), 70, 82.0, 460, 1
    ))

    # Reverse
    route_id_back = f'vilyuysk_yks_{date_str}'
    seg_id_back = f'seg_vilyuysk_yks_{date_str}'

    departure_back = date.replace(hour=16, minute=0)
    arrival_back = date.replace(hour=17, minute=40)
    price_back = vary_price(BASE_PRICES['yakutsk_vilyuysk'] * 1.03)
    insurance_back = price_back * 0.05

    routes.append(generate_route_sql(
        route_id_back, 'vilyuysk', 'yakutsk', departure_back, arrival_back, ns_to_hours(1.67),
        price_back, 82.0, insurance_back, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id_back, route_id_back, 'air', 'Polar Airlines', 'vilyuysk_aprt', 'yakutsk_yks',
        departure_back, arrival_back, price_back, ns_to_hours(1.67), 70, 82.0, 460, 1
    ))

    # ========== Yakutsk ↔ Tiksi (1 flight each direction) ==========
    route_id = f'yks_tiksi_{date_str}'
    seg_id = f'seg_yks_tiksi_{date_str}'

    departure = date.replace(hour=8, minute=30)
    arrival = date.replace(hour=11, minute=0)
    price = vary_price(BASE_PRICES['yakutsk_tiksi'])
    insurance = price * 0.05

    routes.append(generate_route_sql(
        route_id, 'yakutsk', 'tiksi', departure, arrival, ns_to_hours(2.5),
        price, 78.0, insurance, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id, route_id, 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
        departure, arrival, price, ns_to_hours(2.5), 60, 78.0, 1670, 1
    ))

    # Reverse
    route_id_back = f'tiksi_yks_{date_str}'
    seg_id_back = f'seg_tiksi_yks_{date_str}'

    departure_back = date.replace(hour=11, minute=30)
    arrival_back = date.replace(hour=14, minute=0)
    price_back = vary_price(BASE_PRICES['yakutsk_tiksi'] * 1.02)
    insurance_back = price_back * 0.05

    routes.append(generate_route_sql(
        route_id_back, 'tiksi', 'yakutsk', departure_back, arrival_back, ns_to_hours(2.5),
        price_back, 78.0, insurance_back, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id_back, route_id_back, 'air', 'Polar Airlines', 'tiksi_aprt', 'yakutsk_yks',
        departure_back, arrival_back, price_back, ns_to_hours(2.5), 60, 78.0, 1670, 1
    ))

    # ========== Yakutsk ↔ Pokrovsk buses (2 each direction) ==========
    pokrovsk_buses = [('morning', 8), ('afternoon', 14)]

    for time_slot, dep_hour in pokrovsk_buses:
        # Yakutsk → Pokrovsk
        route_id = f'yks_pokrovsk_{date_str}_{time_slot}'
        seg_id = f'seg_yks_pokrovsk_{date_str}_{time_slot[0]}'

        departure = date.replace(hour=dep_hour, minute=0)
        arrival = departure + timedelta(hours=2.5)

        price = vary_price(BASE_PRICES['yakutsk_pokrovsk'])
        insurance = price * 0.05

        routes.append(generate_route_sql(
            route_id, 'yakutsk', 'pokrovsk', departure, arrival, ns_to_hours(2.5),
            price, 86.0, insurance, "['bus']"
        ))

        segments.append(generate_segment_sql(
            seg_id, route_id, 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
            departure, arrival, price, ns_to_hours(2.5), 45, 86.0, 80, 1
        ))

        # Pokrovsk → Yakutsk (3 hours after)
        route_id_back = f'pokrovsk_yks_{date_str}_{time_slot}'
        seg_id_back = f'seg_pokrovsk_yks_{date_str}_{time_slot[0]}'

        departure_back = arrival + timedelta(hours=0.5)
        arrival_back = departure_back + timedelta(hours=2.5)

        routes.append(generate_route_sql(
            route_id_back, 'pokrovsk', 'yakutsk', departure_back, arrival_back, ns_to_hours(2.5),
            price, 86.0, insurance, "['bus']"
        ))

        segments.append(generate_segment_sql(
            seg_id_back, route_id_back, 'bus', 'Avtotrans Yakutia', 'pokrovsk_bus', 'yakutsk_bus',
            departure_back, arrival_back, price, ns_to_hours(2.5), 45, 86.0, 80, 1
        ))

    # ========== Multi-segment: Moscow → Mirny ==========
    route_id = f'msk_mirny_{date_str}'
    seg_id_1 = f'seg_msk_mirny_{date_str}_1'
    seg_id_2 = f'seg_msk_mirny_{date_str}_2'

    dep1 = date.replace(hour=8, minute=0)
    arr1 = date.replace(hour=23, minute=0)
    dep2 = (date + timedelta(days=1)).replace(hour=0, minute=0)
    arr2 = (date + timedelta(days=1)).replace(hour=1, minute=30)

    price1 = vary_price(BASE_PRICES['moscow_yakutsk'])
    price2 = vary_price(BASE_PRICES['yakutsk_mirny'] * 0.9)
    total_price = price1 + price2
    total_insurance = total_price * 0.05

    routes.append(generate_route_sql(
        route_id, 'moscow', 'mirny', dep1, arr2, ns_to_hours(9.5),
        total_price, 89.0, total_insurance, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id_1, route_id, 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
        dep1, arr1, price1, ns_to_hours(8), 180, 92.0, 4884, 1
    ))

    segments.append(generate_segment_sql(
        seg_id_2, route_id, 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
        dep2, arr2, price2, ns_to_hours(1.5), 100, 86.0, 520, 2
    ))

    connections.append(generate_connection_sql(
        route_id, seg_id_1, seg_id_2, ns_to_hours(1), 0, False, ns_to_hours(1), 1
    ))

    # ========== Multi-segment: Mirny → Moscow ==========
    route_id_back = f'mirny_msk_{date_str}'
    seg_id_back_1 = f'seg_mirny_msk_{date_str}_1'
    seg_id_back_2 = f'seg_mirny_msk_{date_str}_2'

    dep1_back = date.replace(hour=11, minute=0)
    arr1_back = date.replace(hour=12, minute=30)
    dep2_back = date.replace(hour=19, minute=0)
    arr2_back = (date + timedelta(days=1)).replace(hour=3, minute=0)

    price1_back = vary_price(BASE_PRICES['yakutsk_mirny'] * 0.95)
    price2_back = vary_price(BASE_PRICES['moscow_yakutsk'] * 1.05)
    total_price_back = price1_back + price2_back
    total_insurance_back = total_price_back * 0.05

    routes.append(generate_route_sql(
        route_id_back, 'mirny', 'moscow', dep1_back, arr2_back, ns_to_hours(10),
        total_price_back, 89.0, total_insurance_back, "['air']"
    ))

    segments.append(generate_segment_sql(
        seg_id_back_1, route_id_back, 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
        dep1_back, arr1_back, price1_back, ns_to_hours(1.5), 100, 88.0, 520, 1
    ))

    segments.append(generate_segment_sql(
        seg_id_back_2, route_id_back, 'air', 'Aeroflot', 'yakutsk_yks', 'moscow_svo',
        dep2_back, arr2_back, price2_back, ns_to_hours(8), 220, 90.0, 4884, 2
    ))

    connections.append(generate_connection_sql(
        route_id_back, seg_id_back_1, seg_id_back_2, ns_to_hours(6.5), 0, False, ns_to_hours(6.5), 1
    ))

    return routes, segments, connections

def main():
    """Generate full week of routes"""
    output_file = 'seed_full_week_routes.sql'

    with open(output_file, 'w', encoding='utf-8') as f:
        f.write("""-- ============================================================================
-- FULL WEEK ROUTES: NOVEMBER 28 - DECEMBER 4, 2025
-- LenaLink - Multi-modal Transport Aggregator for Yakutia
-- Generated by generate_week_routes.py
-- ============================================================================
-- Includes:
-- - Bidirectional routes for all city pairs
-- - Multiple daily flights on major routes
-- - Multi-segment journeys (Moscow → Yakutia cities via Yakutsk)
-- - Price variations across days
-- ============================================================================

BEGIN;

""")

        total_routes = 0
        total_segments = 0
        total_connections = 0

        for day_num in range(DAYS):
            current_date = START_DATE + timedelta(days=day_num)
            date_str = current_date.strftime('%B %d, %Y (%A)')

            f.write(f"""-- ============================================================================
-- DAY {day_num + 1}: {date_str.upper()}
-- ============================================================================

""")

            routes, segments, connections = generate_day_routes(current_date)

            f.write("-- Routes\n")
            for route in routes:
                f.write(route)

            f.write("\n-- Segments\n")
            for segment in segments:
                f.write(segment)

            if connections:
                f.write("\n-- Connections\n")
                for connection in connections:
                    f.write(connection)

            f.write("\n")

            total_routes += len(routes)
            total_segments += len(segments)
            total_connections += len(connections)

        f.write(f"""COMMIT;

-- ============================================================================
-- SUMMARY:
-- - Period: November 28 - December 4, 2025 ({DAYS} days)
-- - Total routes: {total_routes}
-- - Total segments: {total_segments}
-- - Total connections: {total_connections}
--
-- Routes per day (~{total_routes // DAYS}):
-- - Moscow ↔ Yakutsk: 6 flights (3 each direction)
-- - Yakutsk ↔ Mirny: 6 flights (3 each direction)
-- - Yakutsk ↔ Neryungri: 4 flights (2 each direction)
-- - Yakutsk ↔ Udachny: 2 flights (bidirectional)
-- - Yakutsk ↔ Vilyuysk: 2 flights (bidirectional)
-- - Yakutsk ↔ Tiksi: 2 flights (bidirectional)
-- - Yakutsk ↔ Pokrovsk: 4 buses (2 each direction)
-- - Moscow ↔ Mirny: 2 multi-segment (bidirectional)
-- ============================================================================
""")

    print(f"✓ Generated {output_file}")
    print(f"  - {total_routes} routes")
    print(f"  - {total_segments} segments")
    print(f"  - {total_connections} connections")
    print(f"  - Covering {DAYS} days ({START_DATE.strftime('%Y-%m-%d')} to {(START_DATE + timedelta(days=DAYS-1)).strftime('%Y-%m-%d')})")

if __name__ == '__main__':
    main()
