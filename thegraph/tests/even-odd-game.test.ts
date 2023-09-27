import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { BigInt, Address } from "@graphprotocol/graph-ts"
import { BetResult } from "../generated/schema"
import { BetResult as BetResultEvent } from "../generated/EvenOddGame/EvenOddGame"
import { handleBetResult } from "../src/even-odd-game"
import { createBetResultEvent } from "./even-odd-game-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let requestId = BigInt.fromI32(234)
    let amount = BigInt.fromI32(234)
    let callerAddress = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let gameResult = "boolean Not implemented"
    let newBetResultEvent = createBetResultEvent(
      requestId,
      amount,
      callerAddress,
      gameResult
    )
    handleBetResult(newBetResultEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("BetResult created and stored", () => {
    assert.entityCount("BetResult", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "BetResult",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "requestId",
      "234"
    )
    assert.fieldEquals(
      "BetResult",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "amount",
      "234"
    )
    assert.fieldEquals(
      "BetResult",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "callerAddress",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "BetResult",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "gameResult",
      "boolean Not implemented"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
