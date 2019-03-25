import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateComponent = ({ id }) => (
  <div>
    <h1>Create Component</h1>
    <div>
      <label>ID</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>CreatedAt</label>
      <Input placeholder="Time!" />
    </div>
    <div>
      <label>UpdatedAt</label>
      <Input placeholder="Time!" />
    </div>
    <div>
      <label>ComponentType</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>ComponentImplementation</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>LayoutColumn</label>
      <Input placeholder="ID" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateComponent.propTypes = {
  id: PropTypes.string,
}

export default CreateComponent
