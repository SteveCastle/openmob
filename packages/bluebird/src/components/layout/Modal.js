import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  height,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  position,
  flexWrap,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledDimmer = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${fontSize}
  ${display}
  ${position}
  ${flexWrap}
  ${color}
  ${display}
  ${alignItems}
  ${justifyContent}
  z-index: 999;
  top: 0;
  left: 0;
`;

const StyledModal = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${fontSize}
  ${display}
  ${flexWrap}
  ${color}
  ${display}
  ${alignItems}
  ${justifyContent}
  min-width: 600px;
  min-height: 500px;
`;

function Modal({ children, active, onClose }) {
  return active ? (
    <StyledDimmer
      position="absolute"
      width={1}
      height="100%"
      bg={'modalBackground'}
      display="flex"
      justifyContent="center"
      alignItems="center"
      onClick={onClose}
    >
      <StyledModal
        bg="gray.0"
        p={2}
        border={0}
        boxShadow={0}
        borderRadius={2}
        display="flex"
        flexWrap="wrap"
        flexDirection="column"
        fontSize={1}
      >
        {children}
      </StyledModal>
    </StyledDimmer>
  ) : null;
}

Modal.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  active: PropTypes.bool,
  onClose: PropTypes.func
};
Modal.defaultProps = {};

export default Modal;
