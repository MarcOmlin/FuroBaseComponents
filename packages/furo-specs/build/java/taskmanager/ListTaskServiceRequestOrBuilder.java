// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: __bundled/BundledService.proto

package taskmanager;

public interface ListTaskServiceRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:taskmanager.ListTaskServiceRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   *Partial representation (comma separated list of field names), ?fields=
   * </pre>
   *
   * <code>string fields = 1;</code>
   */
  java.lang.String getFields();
  /**
   * <pre>
   *Partial representation (comma separated list of field names), ?fields=
   * </pre>
   *
   * <code>string fields = 1;</code>
   */
  com.google.protobuf.ByteString
      getFieldsBytes();

  /**
   * <pre>
   *The response message will be filtered by the fields before being sent back to the client, filter=[[&amp;#39;id&amp;#39;,&amp;#39;eq&amp;#39;,&amp;#39;1&amp;#39;]]
   * </pre>
   *
   * <code>string filter = 2;</code>
   */
  java.lang.String getFilter();
  /**
   * <pre>
   *The response message will be filtered by the fields before being sent back to the client, filter=[[&amp;#39;id&amp;#39;,&amp;#39;eq&amp;#39;,&amp;#39;1&amp;#39;]]
   * </pre>
   *
   * <code>string filter = 2;</code>
   */
  com.google.protobuf.ByteString
      getFilterBytes();

  /**
   * <pre>
   *Specifies the result ordering for List requests. The default sorting order is ascending, ?order_by=foo desc,bar
   * </pre>
   *
   * <code>string order_by = 3;</code>
   */
  java.lang.String getOrderBy();
  /**
   * <pre>
   *Specifies the result ordering for List requests. The default sorting order is ascending, ?order_by=foo desc,bar
   * </pre>
   *
   * <code>string order_by = 3;</code>
   */
  com.google.protobuf.ByteString
      getOrderByBytes();

  /**
   * <pre>
   *Use this field to specify the maximum number of results to be returned by the server. 
   *The server may further constrain the maximum number of results returned in a single page. 
   *If the page_size is 0, the server will decide the number of results to be returned. page_size=15
   * </pre>
   *
   * <code>string page_size = 4;</code>
   */
  java.lang.String getPageSize();
  /**
   * <pre>
   *Use this field to specify the maximum number of results to be returned by the server. 
   *The server may further constrain the maximum number of results returned in a single page. 
   *If the page_size is 0, the server will decide the number of results to be returned. page_size=15
   * </pre>
   *
   * <code>string page_size = 4;</code>
   */
  com.google.protobuf.ByteString
      getPageSizeBytes();

  /**
   * <pre>
   *Query term to search a {{.name}}
   * </pre>
   *
   * <code>string q = 5;</code>
   */
  java.lang.String getQ();
  /**
   * <pre>
   *Query term to search a {{.name}}
   * </pre>
   *
   * <code>string q = 5;</code>
   */
  com.google.protobuf.ByteString
      getQBytes();

  /**
   * <pre>
   *allows the client to specify which view of the resource it wants to receive in the response. view=BASIC
   * </pre>
   *
   * <code>string view = 6;</code>
   */
  java.lang.String getView();
  /**
   * <pre>
   *allows the client to specify which view of the resource it wants to receive in the response. view=BASIC
   * </pre>
   *
   * <code>string view = 6;</code>
   */
  com.google.protobuf.ByteString
      getViewBytes();
}
