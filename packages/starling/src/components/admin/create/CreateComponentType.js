import React from "react"
import PropTypes from "prop-types"

const CreateComponentType = ({id}) => (
  <div>
    <h1>Create ComponentType</h1>
  </div>
)

CreateComponentType.propTypes = {
  id: PropTypes.string
}

export default CreateComponentType
