defmodule Scratchcards do
  def process_scratchcards(card_strings) do
    cards =
      card_strings
      |> Enum.map(&CardParser.parse_card/1)
      |> IO.inspect()

    total_reward = cards
      |> Enum.map(fn card -> process_card_reward(card, cards) end)
      |> Enum.sum()

    IO.puts(total_reward)
    total_reward
  end

  def process_card_reward(card, all_cards) do
    value =
      card
      |> CardEvaluator.evaluate_card()

    victory_range = (card.id + 1)..(card.id + value)
    IO.inspect(victory_range)

    num_children = Enum.reduce(victory_range, 0, fn n, acc ->
        acc + process_card_reward(Enum.at(all_cards,n), all_cards)
      end)

    IO.puts(num_children)

    num_children + 1
  end

  def process_scratchcards_from_file(input_file) do
    FileReader.read_scratchcard_file(input_file)
    |> process_scratchcards()
  end
end
