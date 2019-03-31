import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateNote = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create Note</h1>
        <Widget>
          <Label>Contact</Label>
          <Input placeholder="ID!" />
        </Widget>
        <Widget>
          <Label>Cause</Label>
          <Input placeholder="ID!" />
        </Widget>
        <Widget>
          <Label>Body</Label>
          <Input placeholder="String" />
        </Widget>

        <Button label="Create" block variant="primary" />
      </Form>
    </Card>
  </Content>
)

CreateNote.propTypes = {
  id: PropTypes.string,
}

export default CreateNote
