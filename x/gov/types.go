package gov

import crypto "github.com/tendermint/go-crypto"

type Vote struct {
	Voter      crypto.address `json:"voter"`       //  address of the voter
	ProposalID int64          `json:"proposal_id"` //  proposalID of the proposal
	Option     string         `json:"option"`      //  option from OptionSet chosen by the voter
}

//-----------------------------------------------------------

// Proposal
type Proposal struct {
	ProposalID   int64     `json:"proposal_id"`   //  ID of the proposal
	Title        string    `json:"title"`         //  Title of the proposal
	Description  string    `json:"description"`   //  Description of the proposal
	ProposalType string    `json:"proposal_type"` //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
	Procedure    Procedure `json:"procedure"`     //  Governance Procedure that the proposal follows proposal

	SubmitBlock  int64     `json:"submit_block"`  //  Height of the block where TxGovSubmitProposal was included
	TotalDeposit sdk.Coins `json:"total_deposit"` //  Current deposit on this proposal. Initial value is set at InitialDeposit
	Deposits     []Deposit `json:"deposits"`      //  Current deposit on this proposal. Initial value is set at InitialDeposit

	VotingStartBlock int64 `json:"voting_start_block"` //  Height of the block where MinDeposit was reached. -1 if MinDeposit is not reached

	ValidatorGovInfos []ValidatorGovInfo `json:"validator_gov_infos"` //  Total voting power when proposal enters voting period (default 0)
	VoteList          []Vote             `json:"vote_list"`           //  Total votes for each option

	TotalVotingPower int64 `json:"total_voting_power"` //  The Total Voting Power
	YesVotes         int64 `json:"yes_votes"`          //  Weight of Yes Votes
	NoVotes          int64 `json:"no_votes"`           //  Weight of No Votes
	NoWithVetoVotes  int64 `json:"no_with_veto_votes"` //  Weight of NoWithVeto Votes
	AbstainVotes     int64 `json:"abstain_votes"`      //  Weight of Abstain Votes
}

func (proposal Proposal) getValidatorGovInfo(validatorAddr crypto.address) ValidatorGovInfo {
	for index, validatorGovInfo := range proposal.ValidatorGovInfos {
		if validatorGovInfo.ValidatorAddr == validatorAddr {
			return validatorGovInfo
		}
	}
	return nil
}

func (proposal Proposal) getVote(voterAddr crypto.address) Vote {
	for index, vote := range proposal.Votes {
		if validatorGovInfo.ValidatorAddr == validatorAddr {
			return validatorGovInfo
		}
	}
	return nil
}

func (proposal Proposal) isActive() bool {
	return VotingStartBlock >= 0
}

func (proposal Proposal) updateTally(option string, amount int64) {
	switch option {
	case "Yes":
		proposal.YesVotes += votingPower
	case "No":
		proposal.NoVotes += votingPower
	case "NoWithVeto":
		proposal.NoWithVetoVotes += votingPower
	case "Abstain":
		proposal.AbstainVotes += votingPower
	}
}

// Procedure
type Procedure struct {
	VotingPeriod      int64             `json:"voting_period"`      //  Length of the voting period. Initial value: 2 weeks
	MinDeposit        sdk.Coins         `json:"min_deposit"`        //  Minimum deposit for a proposal to enter voting period.
	ProposalTypes     []string          `json:"proposal_type"`      //  Types available to submitters. {PlainTextProposal, SoftwareUpgradeProposal}
	Threshold         rational.Rational `json:"threshold"`          //  Minimum propotion of Yes votes for proposal to pass. Initial value: 0.5
	Veto              rational.Rational `json:"veto"`               //  Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Initial value: 1/3
	MaxDepositPeriod  int64             `json:"max_deposit_period"` //  Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months
	GovernancePenalty int64             `json:"governance_penalty"` //  Penalty if validator does not vote
}

// Deposit
type Deposit struct {
	Depositer crypto.address `json:"depositer"` //  Address of the depositer
	Amount    sdk.Coins      `json:"amount"`    //  Deposit amount
}

type ValidatorGovInfo struct {
	ProposalID      int64          //  Id of the Proposal this validator
	ValidatorAddr   crypto.address //  Address of the validator
	InitVotingPower int64          //  Voting power of validator when proposal enters voting period
	Minus           int64          //  Minus of validator, used to compute validator's voting power
	LastVoteWeight  int64          //  Weight of the last vote by validator at time of casting, -1 if hasn't voted yet
}

type ProposalQueue []int64