import re
import numpy as np
from scipy.optimize import milp, LinearConstraint, Bounds


class JMachine:
    def __init__(self, buttons, joltages):
        self.buttons = buttons
        self.joltages = joltages


def part2(lines):
    """Solve part 2 for non-square matrices"""
    jmachines = parse_jmachines(lines)
    total_presses = 0
    for j in jmachines:
        presses = min_jpresses(j)
        total_presses += presses
    return total_presses


def min_jpresses(jmachine):
    """
    Find minimum button presses to achieve target joltages.
    Works with non-square matrices (different number of buttons and joltages).
    All solutions must be integers (no half button presses allowed).
    """
    # Convert buttons to float matrix for optimization
    button_matrix = np.array(jmachine.buttons, dtype=float).T  # Transpose to get (n_joltages x n_buttons)
    target = np.array(jmachine.joltages, dtype=float)
    
    n_joltages = len(target)
    n_buttons = len(jmachine.buttons)
    
    print(f"Matrix dimensions: {n_joltages} joltages x {n_buttons} buttons")
    
    try:
        vars_result, obj_value = solve_min_positive_coeffs(button_matrix, target)
        if vars_result is not None:
            print(f"Button presses (integers): {vars_result}")
            print(f"Total presses: {obj_value}")
            return obj_value  # Already an integer
        else:
            print("No integer solution found")
            return 0
    except Exception as e:
        print(f"Error: {e}")
        return 0


def solve_min_positive_coeffs(vectors_matrix, target, tol=1e-6):
    """
    Solve for minimum positive INTEGER coefficients using mixed-integer linear programming.
    
    Args:
        vectors_matrix: numpy array of shape (n_joltages, n_buttons)
                       Each column is a button's effect on joltages
        target: numpy array of target joltages (length n_joltages)
        tol: tolerance for solution validation
    
    Returns:
        (coefficients, objective_value) or (None, None) if infeasible
    """
    n_joltages, n_buttons = vectors_matrix.shape
    
    if n_joltages == 0 or n_buttons == 0:
        raise ValueError("Matrix dimensions must be non-zero")
    
    # Check for all-zero columns (buttons with no effect)
    for j in range(n_buttons):
        if np.all(vectors_matrix[:, j] == 0):
            raise ValueError(f"Button {j} has no effect (all zeros)")
    
    # Check for all-zero rows (joltages not affected by any button)
    for i in range(n_joltages):
        if np.all(vectors_matrix[i, :] == 0):
            raise ValueError(f"Joltage {i} is not affected by any button")
    
    # Mixed-Integer Linear Programming formulation:
    # minimize: sum(x_i) for i in 0..n_buttons-1
    # subject to: A @ x = b (equality constraints)
    #             x >= 0 (bounds)
    #             x must be integers
    
    # Objective: minimize sum of all button presses
    c = np.ones(n_buttons)
    
    # Equality constraint: button_matrix @ x = target
    # Using LinearConstraint with lb = ub = target for equality
    constraints = LinearConstraint(vectors_matrix, lb=target, ub=target)
    
    # Bounds: x_i >= 0 for all buttons (no upper bound specified)
    bounds = Bounds(lb=np.zeros(n_buttons), ub=np.full(n_buttons, np.inf))
    
    # All variables must be integers
    integrality = np.ones(n_buttons)  # 1 = integer, 0 = continuous
    
    # Solve using scipy's MILP solver
    result = milp(
        c=c,
        constraints=constraints,
        bounds=bounds,
        integrality=integrality,
        options={'disp': False}
    )
    
    if result.success:
        # Validate solution - all values should be integers
        if not np.allclose(result.x, np.round(result.x), atol=1e-9):
            print(f"Warning: Solution contains non-integer values: {result.x}")
            return None, None
        
        # Round to nearest integer (should already be very close)
        int_solution = np.round(result.x).astype(int)
        
        # Validate that the solution achieves the target
        achieved = vectors_matrix @ int_solution
        if np.allclose(achieved, target, atol=tol):
            return int_solution, int(np.sum(int_solution))
        else:
            print(f"Solution doesn't match target within tolerance")
            print(f"Target:   {target}")
            print(f"Achieved: {achieved}")
            return None, None
    else:
        print(f"Optimization failed: {result.message}")
        return None, None


def parse_jmachines(lines):
    """Parse all machine configurations from input lines"""
    machines = []
    for line in lines:
        if line.strip():
            machines.append(parse_jmachine(line))
    return machines


def parse_jmachine(line):
    """
    Parse a single machine configuration line.
    
    Format: [###...] (button1_indices) (button2_indices) ... {joltage1,joltage2,...}
    Example: [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
    """
    # Extract joltages from {n,n,n,...}
    joltages_match = re.search(r'\{([^}]+)\}', line)
    if not joltages_match:
        raise ValueError(f"No joltages found in line: {line}")
    
    joltages_str = joltages_match.group(1)
    joltages = [int(x.strip()) for x in joltages_str.split(',')]
    matrix_length = len(joltages)
    
    # Extract all button patterns (n) or (n,m,...)
    buttons_matches = re.findall(r'\(([^)]+)\)', line)
    if not buttons_matches:
        raise ValueError(f"No buttons found in line: {line}")
    
    buttons = []
    for button_str in buttons_matches:
        buttons.append(parse_matrix_button(button_str, matrix_length))
    
    return JMachine(buttons, joltages)


def parse_matrix_button(button_string, length):
    """
    Parse button indices and create binary vector.
    
    Args:
        button_string: comma-separated indices like "1,3" or single "2"
        length: total length of the vector (number of joltages)
    
    Returns:
        List of 1s and 0s indicating which joltages this button affects
    """
    if not button_string.strip():
        return [0] * length
    
    indices = [int(x.strip()) for x in button_string.split(',')]
    
    output = []
    for i in range(length):
        if i in indices:
            output.append(1)
        else:
            output.append(0)
    
    return output


def main():
    """Main entry point"""
    # Read input file
    with open('input.txt', 'r') as f:
        lines = f.readlines()
    
    result = part2(lines)
    print(f"\nFinal answer: {result}")


if __name__ == "__main__":
    main()
