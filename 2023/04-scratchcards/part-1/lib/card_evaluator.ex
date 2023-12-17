defmodule CardEvaluator do
  def evaluate_card(card) do
    IO.puts("Evaluating Card")
    IO.inspect(card)

    num_matches = card.owned_numbers
    |> Enum.filter(&Enum.member?(card.result_numbers, &1))
    |> IO.inspect()
    |> Enum.count()

    value_of_card = if (num_matches > 0) do
      :math.pow(2, num_matches - 1)
    else
      0
    end

    IO.puts("Number of Winning items: #{num_matches}")
    IO.puts("Value: #{value_of_card}")

    value_of_card
  end
end
