defmodule Card do
  defstruct id: nil, result_numbers: [], owned_numbers: []
end

defmodule CardParser do
  def parse_card(card_string) do

    [card_id_string, data_string] =
      card_string
      |> String.split(":")
      |> Enum.map(&String.trim/1)

    card_id =
      card_id_string
      |> String.split(~r/\s+/)
      |> List.last()
      |> String.to_integer()

    [results_s, owned_s] =
      data_string
    |> String.split("|")
    |> Enum.map(&String.trim/1)
    |> Enum.map(&String.split(&1, ~r/\s+/))

    results = results_s
    |> Enum.map(&String.to_integer/1)

    owned = owned_s
    |> Enum.map(&String.to_integer/1)

   %Card{id: card_id, result_numbers: results, owned_numbers: owned}
  end
end
