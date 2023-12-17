defmodule Scratchcards do
  def process_scratchcards(card_strings) do
    cards =
      card_strings
      |> Enum.map(&CardParser.parse_card/1)

    total_reward = cards
      |> Enum.map(fn card -> process_card_reward(card, cards) end)
      |> Enum.sum()

    total_reward
  end

  def process_card_reward(card, all_cards) do
    value =
      card
      |> CardEvaluator.evaluate_card()

    total = if (value > 0) do
      victory_range = (card.id + 1)..(card.id + value)

      num_children = Enum.reduce(victory_range, 0, fn n, acc ->
          acc + process_card_reward(Enum.at(all_cards,n-1), all_cards)
        end)
      num_children + 1
    else
      1
    end

    total
  end

  def process_scratchcards_from_file(input_file) do
    res = FileReader.read_scratchcard_file(input_file)
    |> process_scratchcards()
    IO.puts("Answer: #{res}")
  end
end
