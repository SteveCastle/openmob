import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateLayoutRow = ({ id }) => (
  <div>
    <h1>Create LayoutRow</h1>
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
      <label>Layout</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>Container</label>
      <Input placeholder="Boolean" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateLayoutRow.propTypes = {
  id: PropTypes.string,
}

export default CreateLayoutRow
