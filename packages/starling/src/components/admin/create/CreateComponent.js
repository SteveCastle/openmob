import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateComponent = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create Component</h1>
        <Widget>
          <Label>ID</Label>
          <Input placeholder="ID!" />
        </Widget>
        <Widget>
          <Label>CreatedAt</Label>
          <Input placeholder="Time!" />
        </Widget>
        <Widget>
          <Label>UpdatedAt</Label>
          <Input placeholder="Time!" />
        </Widget>
        <Widget>
          <Label>ComponentType</Label>
          <Input placeholder="ID!" />
        </Widget>
        <Widget>
          <Label>ComponentImplementation</Label>
          <Input placeholder="ID!" />
        </Widget>
        <Widget>
          <Label>LayoutColumn</Label>
          <Input placeholder="ID" />
        </Widget>

        <Button label="Create" block variant="primary" />
      </Form>
    </Card>
  </Content>
)

CreateComponent.propTypes = {
  id: PropTypes.string,
}

export default CreateComponent
