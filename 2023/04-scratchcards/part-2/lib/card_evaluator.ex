defmodule CardEvaluator do
  def evaluate_card(card) do
    num_matches = card.owned_numbers
    |> Enum.filter(&Enum.member?(card.result_numbers, &1))
    |> Enum.count()

    IO.puts("Value of card ##{card.id} is #{num_matches}")
    num_matches
  end
end
