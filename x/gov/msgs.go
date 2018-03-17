package gov

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//-----------------------------------------------------------
// SubmitProposalMsg

type SubmitProposalMsg struct {
	Title          string      //  Title of the proposal
	Description    string      //  Description of the proposal
	ProposalType   string      //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
	Proposer       sdk.Address //  Address of the proposer
	InitialDeposit sdk.Coins   //  Initial deposit paid by sender. Must be strictly positive.
}

func NewSubmitProposalMsg(title string, description string, proposalType string, proposer sdk.Address, initialDeposit int64) SubmitProposalMsg {
	return SubmitProposalMsg{
		Title:          title,
		Description:    description,
		ProposalType:   proposalType,
		Proposer:       proposer,
		InitialDeposit: initialDeposit,
	}
}

// Implements Msg.
func (msg SubmitProposalMsg) Type() string { return "gov" }

// Implements Msg.
func (msg SubmitProposalMsg) ValidateBasic() sdk.Error {

	if len(msg.Title) == 0 {
		return ErrInvalidAttribute(msg.Title) // TODO: Proper Error
	}
	if len(msg.Description) == 0 {
		return ErrInvalidAttribute(msg.Description) // TODO: Proper Error
	}
	if len(msg.ProposalType) == 0 {
		return ErrInvalidAttribute(msg.ProposalType) // TODO: Proper Error
	}
	if len(msg.Proposer) == 0 {
		return ErrInvalidAddress(msg.Proposer.String())
	}
	if !msg.Amount.IsValid() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	if !msg.Amount.IsPositive() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	return nil
}

func (msg SubmitProposalMsg) String() string {
	return fmt.Sprintf("SubmitProposalMsg{%v, %v, %v, %v}", msg.Title, msg.Description, msg.ProposalType, msg.InitialDeposit)
}

// Implements Msg.
func (msg SubmitProposalMsg) Get(key interface{}) (value interface{}) {
	return nil
}

// Implements Msg.
func (msg SubmitProposalMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	return b
}

// Implements Msg.
func (msg SubmitProposalMsg) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Proposer}
}

//-----------------------------------------------------------
// DepositMsg

type DepositMsg struct {
	ProposalID int64       `json:"proposal_id"` // ID of the proposal
	Depositer  sdk.Address `json:"depositer"`   // Address of the depositer
	Amount     sdk.Coins   `json:"amount"`      // Coins to add to the proposal's deposit
}

func NewDepositMsgMsg(proposalID int64, depositer sdk.Address, amount sdk.Coins) DepositMsg {
	return DepositMsg{
		ProposalID: proposalID,
		Depositer:  depositer,
		Amount:     Amount,
	}
}

// Implements Msg.
func (msg DepositMsg) Type() string { return "gov" }

// Implements Msg.
func (msg DepositMsg) ValidateBasic() sdk.Error {
	if len(msg.Depositer) == 0 {
		return ErrInvalidAddress(msg.Depositer.String())
	}
	if !msg.Amount.IsValid() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	if !msg.Amount.IsPositive() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	return nil
}

func (msg DepositMsg) String() string {
	return fmt.Sprintf("DepositMsg{%v=>%v: %v}", msg.Depositer, msg.ProposerId, msg.Amount)
}

// Implements Msg.
func (msg DepositMsg) Get(key interface{}) (value interface{}) {
	return nil
}

// Implements Msg.
func (msg DepositMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	return b
}

// Implements Msg.
func (msg DepositMsg) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Depositer}
}

//-----------------------------------------------------------
// VoteMsg

type VoteMsg struct {
	Voter      sdk.Address //  address of the voter
	ProposalID int64       //  proposalID of the proposal
	Option     string      //  option from OptionSet chosen by the voter
}

func NewVoteMsg(voter sdk.Address, proposalID int64, option string) VoteMsg {
	return VoteMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Option:     option,
	}
}

// Implements Msg.
func (msg VoteMsg) Type() string { return "gov" }

// Implements Msg.
func (msg VoteMsg) ValidateBasic() sdk.Error {

	if len(msg.Voter) == 0 {
		return ErrInvalidAddress(msg.Voter.String())
	}
	if msg.Option != "Yes" || msg.Option != "No" || msg.Option != "NoWithVeto" || msg.Option != "Abstain" {
		return ErrInvalidAttribute(msg.Option) // TODO: Proper Error
	}
	return nil
}

func (msg VoteMsg) String() string {
	return fmt.Sprintf("VoteMsg{%v - %v}", msg.ProposalID, msg.Option)
}

// Implements Msg.
func (msg VoteMsg) Get(key interface{}) (value interface{}) {
	return nil
}

// Implements Msg.
func (msg VoteMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	return b
}

// Implements Msg.
func (msg VoteMsg) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Voter}
}
