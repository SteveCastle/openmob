import React from "react"
import PropTypes from "prop-types"

const CreateComponent = ({id}) => (
  <div>
    <h1>Create Component</h1>
  </div>
)

CreateComponent.propTypes = {
  id: PropTypes.string
}

export default CreateComponent
