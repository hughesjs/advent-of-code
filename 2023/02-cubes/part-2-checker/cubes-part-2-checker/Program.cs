using System.Diagnostics;
using System.Text.RegularExpressions;
// 




string[] inputLines = File.ReadAllLines("./input-data").Where(l => !l.Contains("Answer")).ToArray();
Regex lineRegex = new(@"^Game\s(?<GameId>\d+):\s(?<GameData>.*)$");
Regex revealRegex = new(@"^(?<Val>\d*)\s(?<Colour>red|green|blue)$");
Regex resultRegex = new(@"^GameId:\s(?<GameId>\d+),\sR:\s(?<R>\d+),\sG:\s(?<G>\d+),\sB:\s(?<B>\d+),\sPOW:\s(?<POW>\d+)$");

List<PowerStats> powerStats = inputLines.Select(GetPower).ToList();

List<PowerStats> passedIn = new();
while (Console.ReadLine() is { } lineIn)
{
    if (lineIn.Contains("Answer")) continue;
    Match match = resultRegex.Match(lineIn);
    PowerStats stat = new()
    {
        GameId = int.Parse(match.Groups["GameId"].Value),
        R = int.Parse(match.Groups["R"].Value),
        G = int.Parse(match.Groups["G"].Value),
        B = int.Parse(match.Groups["B"].Value),
        Pow = int.Parse(match.Groups["POW"].Value)
    };

    passedIn.Add(stat);
}

for (int i = 0; i < powerStats.Count; ++i)
{
    if (powerStats[i] != passedIn[i])
    {
        Console.WriteLine($"Correct: {powerStats[i]}");
        Console.WriteLine($"Borked: {passedIn[i]}");
    }
}


PowerStats GetPower(string s)
{
    Match match = lineRegex.Match(s);
    int maxR = 0 , maxG = 0 , maxB = 0;
    foreach (string round in match.Groups["GameData"].Value.Split(';'))
    {
        foreach (string reveal in round.Split(',').Select(r => r.Trim()))
        {
            Match revMatch = revealRegex.Match(reveal);
            int val = int.Parse(revMatch.Groups["Val"].Value);
            string color = revMatch.Groups["Colour"].Value;
            switch (color)
            {
                case "red":
                {
                    maxR = Math.Max(maxR, val);
                    break;
                }
                case "green":
                {
                    maxG = Math.Max(maxG, val);
                    break;
                }
                case "blue":
                {
                    maxB = Math.Max(maxB, val);
                    break;
                }
            }
        }
    }


    PowerStats stat = new()
    {
        GameId = int.Parse(match.Groups["GameId"].Value),
        R = maxR,
        G = maxG,
        B = maxB,
        Pow = maxR * maxG * maxB
    };

    return stat;
}

record PowerStats
{
    public int GameId  { get; init; }
    public int R  { get; init;  }
    public int G  { get; init; }
    public int B { get; init;  }

    public int Pow { get; init; }

    public override string ToString()
    {
        return $"GameId: {GameId} R: {R}, G: {G}, B: {B}, POW: {Pow}";
    }
}