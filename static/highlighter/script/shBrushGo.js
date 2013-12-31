/** vim: set noexpandtab tabstop=8 shiftwidth=8 softtabstop=8:
 * SyntaxHighlighter
 * http://alexgorbatchev.com/SyntaxHighlighter
 *
 * SyntaxHighlighter is donationware. If you are using it, please donate.
 * http://alexgorbatchev.com/SyntaxHighlighter/donate.html
 *
 * @version
 * 3.0.83 (July 02 2010)
 * 
 * @copyright
 * Copyright (C) 2011 Christopher Dunn.
 * (based on shBrushCpp.js by Alex Gorbatchev)
 *
 * @license
 * Dual licensed under the MIT and GPL licenses.
 */
;(function()
{
	// CommonJS
	typeof(require) != 'undefined' ? SyntaxHighlighter = require('shCore').SyntaxHighlighter : null;

	function Brush()
	{
		// Copyright 2006 Shin, YoungJin
	
		var datatypes =	'int int8 int16 int32 int64 ' +
				'uint uint8 uint16 unint32 uint64 uintptr ' +
				'float32 float64 ' +
				'complex64 complex128 ' +
				'byte string ' +
				'chan map ' +
				'struct interface ';

		var keywords =	'func var type package import const ' +
				'for if while else range ' +
				'select switch case default';
					
		var keywords_bold = 'go defer goto break continue return fallthrough';
		var constants = 'nil true false iota _';
		var functions =	'append cap close complex copy imag len ' +
				'make new panic print println real recover';

		this.regexList = [
			{ regex: SyntaxHighlighter.regexLib.singleLineCComments,	css: 'comments' },			// one line comments
			{ regex: SyntaxHighlighter.regexLib.multiLineCComments,		css: 'comments' },			// multiline comments
			{ regex: SyntaxHighlighter.regexLib.doubleQuotedString,		css: 'string' },			// strings
			{ regex: SyntaxHighlighter.regexLib.singleQuotedString,		css: 'string bold' },			// strings
			{ regex: /^ *#.*/gm,						css: 'preprocessor' },
			{ regex: /[a-zA-Z_]+(?=\()/gm,					css: 'functions' },
			{ regex: /\b[0-9]+\b/gm,					css: 'constants' },
			{ regex: new RegExp(this.getKeywords(datatypes), 'gm'),		css: 'color1 bold' },
			{ regex: new RegExp(this.getKeywords(constants), 'gm'),		css: 'constants' },
			{ regex: new RegExp(this.getKeywords(functions), 'gm'),		css: 'functions bold' },
			{ regex: new RegExp(this.getKeywords(keywords), 'gm'),		css: 'keyword' },
			{ regex: new RegExp(this.getKeywords(keywords_bold), 'gm'),	css: 'keyword bold' }
			];
	};

	Brush.prototype	= new SyntaxHighlighter.Highlighter();
	Brush.aliases	= ['go', 'golang'];

	SyntaxHighlighter.brushes.Go = Brush;

	// CommonJS
	typeof(exports) != 'undefined' ? exports.Brush = Brush : null;
})();