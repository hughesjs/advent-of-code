using Xunit;

namespace SecretEntrance.Tests;

public class SolverTests
{
    [Fact]
    public void Solve_WithTestData_Returns6()
    {
        var testData = new[]
        {
            "L68",
            "L30",
            "R48",
            "L5",
            "R60",
            "L55",
            "L1",
            "L99",
            "R14",
            "L82"
        };

        var result = Solver.Solve(testData);

        Assert.Equal(6, result);
    }
}
