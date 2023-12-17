defmodule Scratchcards do
  def process_scratchcards(card_data) do
    card_data
    |> Enum.map(&CardParser.parse_card/1)
    |> Enum.map(&CardEvaluator.evaluate_card/1)
    |> Enum.sum()
  end


  def process_scratchcards_from_file(input_file) do
    FileReader.read_scratchcard_file(input_file)
    |> Enum.map(&process_scratchcards/1)
  end
end
