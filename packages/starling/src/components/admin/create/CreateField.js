import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateField = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create Field</h1>
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
          <Label>FieldType</Label>
          <Input placeholder="ID!" />
        </Widget>
        <Widget>
          <Label>StringValue</Label>
          <Input placeholder="String" />
        </Widget>
        <Widget>
          <Label>IntValue</Label>
          <Input placeholder="Int" />
        </Widget>
        <Widget>
          <Label>FloatValue</Label>
          <Input placeholder="Float" />
        </Widget>
        <Widget>
          <Label>BooleanValue</Label>
          <Input placeholder="Boolean" />
        </Widget>
        <Widget>
          <Label>DateTimeValue</Label>
          <Input placeholder="Time" />
        </Widget>
        <Widget>
          <Label>Component</Label>
          <Input placeholder="ID" />
        </Widget>

        <Button label="Edit" />
      </Form>
    </Card>
  </Content>
)

CreateField.propTypes = {
  id: PropTypes.string,
}

export default CreateField
