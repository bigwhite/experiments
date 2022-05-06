// Generated from CSV.g4 by ANTLR 4.10.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link CSVParser}.
 */
public interface CSVListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link CSVParser#csvFile}.
	 * @param ctx the parse tree
	 */
	void enterCsvFile(CSVParser.CsvFileContext ctx);
	/**
	 * Exit a parse tree produced by {@link CSVParser#csvFile}.
	 * @param ctx the parse tree
	 */
	void exitCsvFile(CSVParser.CsvFileContext ctx);
	/**
	 * Enter a parse tree produced by {@link CSVParser#hdr}.
	 * @param ctx the parse tree
	 */
	void enterHdr(CSVParser.HdrContext ctx);
	/**
	 * Exit a parse tree produced by {@link CSVParser#hdr}.
	 * @param ctx the parse tree
	 */
	void exitHdr(CSVParser.HdrContext ctx);
	/**
	 * Enter a parse tree produced by {@link CSVParser#row}.
	 * @param ctx the parse tree
	 */
	void enterRow(CSVParser.RowContext ctx);
	/**
	 * Exit a parse tree produced by {@link CSVParser#row}.
	 * @param ctx the parse tree
	 */
	void exitRow(CSVParser.RowContext ctx);
	/**
	 * Enter a parse tree produced by {@link CSVParser#field}.
	 * @param ctx the parse tree
	 */
	void enterField(CSVParser.FieldContext ctx);
	/**
	 * Exit a parse tree produced by {@link CSVParser#field}.
	 * @param ctx the parse tree
	 */
	void exitField(CSVParser.FieldContext ctx);
}