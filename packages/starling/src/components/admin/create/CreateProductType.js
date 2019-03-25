import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateProductType = ({ id }) => (
  <div>
    <h1>Create ProductType</h1>
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
      <label>Title</label>
      <Input placeholder="String!" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateProductType.propTypes = {
  id: PropTypes.string,
}

export default CreateProductType
