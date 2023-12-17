defmodule FileReader do
  def read_scratchcard_file(filePath) do
    File.stream!(filePath)
    |> Enum.map(&String.trim/1)
  end
end
