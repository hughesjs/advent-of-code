defmodule CardEvaluator do
  def evaluate_card(card) do
    num_matches = card.owned_numbers
    |> Enum.filter(&Enum.member?(card.result_numbers, &1))
    |> Enum.count()

    value_of_card = if (num_matches > 0) do
      :math.pow(2, num_matches - 1)
    else
      0
    end

    value_of_card
  end
end
