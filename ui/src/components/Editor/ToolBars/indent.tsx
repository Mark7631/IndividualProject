import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const Indent = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'indent',
    tip: t('indent.text'),
  };
  const handleClick = (ctx) => {
    context = ctx;
    const { editor, replaceLines } = context;

    replaceLines((line) => {
      line = `    ${line}`;
      return line;
    });
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Indent);
