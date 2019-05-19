import React from 'react';
import { shallow } from 'enzyme';
import SimpleContent from '../SimpleContent';

it('renders without crashing', () => {
  shallow(<SimpleContent />);
});
