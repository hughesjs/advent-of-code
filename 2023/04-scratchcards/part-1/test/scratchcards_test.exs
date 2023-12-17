defmodule ScratchcardsTest do
  use ExUnit.Case
  doctest Scratchcards

  test "greets the world" do
    assert Scratchcards.hello() == :world
  end
end
