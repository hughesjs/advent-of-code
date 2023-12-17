defmodule Scratchcards do
  def process_scratchcards(card_data) do
    total_val = card_data
    |> Enum.map(&CardParser.parse_card/1)
    |> Enum.map(&CardEvaluator.evaluate_card/1)
    |> Enum.sum()

    IO.puts("Answer: #{total_val}")
    total_val
  end


  def process_scratchcards_from_file(input_file) do
    FileReader.read_scratchcard_file(input_file)
    |> process_scratchcards()
  end
end
