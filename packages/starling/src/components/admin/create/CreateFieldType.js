import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateFieldType = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create FieldType</h1>
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
          <Label>Title</Label>
          <Input placeholder="String!" />
        </Widget>
        <Widget>
          <Label>DataType</Label>
          <Input placeholder="String!" />
        </Widget>
        <Widget>
          <Label>PropName</Label>
          <Input placeholder="String!" />
        </Widget>
        <Widget>
          <Label>StringValueDefault</Label>
          <Input placeholder="String" />
        </Widget>
        <Widget>
          <Label>IntValueDefault</Label>
          <Input placeholder="Int" />
        </Widget>
        <Widget>
          <Label>FloatValueDefault</Label>
          <Input placeholder="Float" />
        </Widget>
        <Widget>
          <Label>BooleanValueDefault</Label>
          <Input placeholder="Boolean" />
        </Widget>
        <Widget>
          <Label>DateTimeValueDefault</Label>
          <Input placeholder="Time" />
        </Widget>
        <Widget>
          <Label>ComponentType</Label>
          <Input placeholder="ID" />
        </Widget>

        <Button label="Create" block />
      </Form>
    </Card>
  </Content>
)

CreateFieldType.propTypes = {
  id: PropTypes.string,
}

export default CreateFieldType
