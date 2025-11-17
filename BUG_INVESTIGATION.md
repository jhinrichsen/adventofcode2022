# Bug Investigation: Days 20 and 24

## Day 20: Grove Positioning System

### Tries Table
| Attempt | Value    | Status | Notes |
|---------|----------|--------|-------|
| 1       | -6459    | Wrong  | Too high (commented in test) |
| 2       | -6750    | Wrong  | Current value, needs investigation |

### Example Status
- Example test: PASSES (expected: 3)
- Example validates algorithm is fundamentally correct

### Current Investigation
Need to verify:
1. Is the mixing algorithm correct?
2. Are we handling modular arithmetic properly for wrapping?
3. Is the zero-finding and coordinate calculation correct?

---

## Day 24: Blizzard Basin

### Tries Table
| Attempt | Value | Status | Notes |
|---------|-------|--------|-------|
| 1       | 300   | Wrong  | Current value, needs investigation |

### Example Status
- Example test: PASSES (expected: 18)
- Example validates pathfinding algorithm is correct

### Current Investigation
Need to verify:
1. Are we counting minutes correctly (off-by-one)?
2. Is the starting position handling correct?
3. Are blizzard positions calculated correctly at each timestep?
4. Is the wrapping logic for blizzards correct?

---

## Next Steps
1. Re-read problem statements carefully
2. Debug with smaller test cases
3. Check edge cases and boundary conditions
4. Verify all intermediate calculations
