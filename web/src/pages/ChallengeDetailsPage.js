import * as React from "react"
import { connect } from "react-redux"
import PropTypes from "prop-types"

import {
  fetchChallengeDetail as fetchChallengeDetailAction,
  buyChallenge as buyChallengeAction,
  validateChallenge as validateChallengeAction,
  closeChallenge as closeChallengeAction,
} from "../actions/seasons"
import ChallengeBuyButton from "../components/challenges/ChallengeBuyButton"
import ChallengeValidateButton from "../components/challenges/ChallengeValidateButton"
import ChallengeCloseButton from "../components/challenges/ChallengeCloseButton"
import ValidationsList from "../components/challenges/ValidationsList"

import styles from "./styles/ChallengeDetailsPage.module.css"

import { Page, Grid, Dimmer, Button } from "tabler-react"

class ChallengeDetailsPage extends React.PureComponent {
  componentDidMount() {
    const { fetchChallengeDetailAction, uri } = this.props
    const challengeID = uri.split("/")[3]
    fetchChallengeDetailAction(challengeID)
  }

  render() {
    const {
      challenge,
      activeTeam: { id: teamID } = { id: "no id" },
      buyChallengeAction,
      validateChallengeAction,
      closeChallengeAction,
    } = this.props

    const {
      flavor: { challenge: flavorChallenge } = { challenge: "no challenge" },
      subscriptions,
    } = challenge || {}

    if (!challenge) {
      return <Dimmer active />
    }
    const subscription =
      subscriptions &&
      challenge.subscriptions.find(item => item.status === "Active")

    const validations = subscription && subscription.validations

    return (
      <Page.Content title={flavorChallenge.name}>
        <Grid.Row>
          <Grid.Col lg={4} md={4} sm={4} xs={4}>
            <h4>Author</h4>
            <p className={styles.p}>{flavorChallenge.author}</p>
          </Grid.Col>
          <Grid.Col lg={4} md={4} sm={4} xs={4}>
            <h4>Page</h4>
            <Button
              href={flavorChallenge.homepage}
              target="_blank"
              RootComponent="a"
              social="github"
            >
              Visit page
            </Button>
          </Grid.Col>
          <Grid.Col lg={4} md={4} sm={4} xs={4}>
            <h4>Actions</h4>
            <Button.List>
              <ChallengeBuyButton
                challenge={challenge}
                buyChallenge={buyChallengeAction}
                teamID={teamID}
              />
              <ChallengeCloseButton
                challenge={challenge}
                closeChallenge={closeChallengeAction}
              />
            </Button.List>
          </Grid.Col>
        </Grid.Row>
        <hr />
        <Grid.Row cards={true}>
          {subscriptions && (
            <Grid.Col lg={12} md={12} sm={12} xs={12}>
              <div style={{ marginBottom: "1rem" }}>
                <h3>Validations</h3>
                <Button.List>
                  <ChallengeValidateButton
                    challenge={challenge}
                    validateChallenge={validateChallengeAction}
                  />
                </Button.List>
              </div>
              {validations && <ValidationsList validations={validations} />}
            </Grid.Col>
          )}
        </Grid.Row>
      </Page.Content>
    )
  }
}

ChallengeDetailsPage.propTypes = {
  fetchChallengeDetailAction: PropTypes.func,
}

const mapStateToProps = state => ({
  challenge: state.seasons.challengeInDetail,
  activeTeam: state.seasons.activeTeam,
})

const mapDispatchToProps = {
  buyChallengeAction: (challengeID, teamID, seasonId) =>
    buyChallengeAction(challengeID, teamID, seasonId),
  validateChallengeAction: (validationData, seasonId) =>
    validateChallengeAction(validationData, seasonId),
  closeChallengeAction: subscriptionID => closeChallengeAction(subscriptionID),
  fetchChallengeDetailAction: challengeID =>
    fetchChallengeDetailAction(challengeID),
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ChallengeDetailsPage)
