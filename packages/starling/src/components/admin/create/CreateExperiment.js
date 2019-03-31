import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateExperiment = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create Experiment</h1>
        <Widget>
          <Label>Title</Label>
          <Input placeholder="String!" />
        </Widget>
        <Widget>
          <Label>LandingPage</Label>
          <Input placeholder="ID" />
        </Widget>

        <Button label="Create" block variant="primary" />
      </Form>
    </Card>
  </Content>
)

CreateExperiment.propTypes = {
  id: PropTypes.string,
}

export default CreateExperiment
