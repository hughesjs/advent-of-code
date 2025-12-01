var lines = File.ReadLines("./input-data");
Console.WriteLine($"Answer: {Solver.Solve(lines)}");

public static class Solver
{
    private const int StartVal = 50;

    public static int Solve(IEnumerable<string> lines)
    {
        int dial = StartVal;
        int acc = 0;
        foreach (var line in lines)
        {
            int multiplier = line[0] == 'R' ? 1 : -1;
            int num = int.Parse(line[1..]);
            dial = (dial + (multiplier * num)) % 100;
            if (dial == 0)
            {
                acc++;
            }
        }
        return acc;
    }
}
