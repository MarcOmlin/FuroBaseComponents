// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: __bundled/BundledService.proto

package taskmanager;

/**
 * Protobuf type {@code taskmanager.ListProjectMembersServiceRequest}
 */
public  final class ListProjectMembersServiceRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:taskmanager.ListProjectMembersServiceRequest)
    ListProjectMembersServiceRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ListProjectMembersServiceRequest.newBuilder() to construct.
  private ListProjectMembersServiceRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ListProjectMembersServiceRequest() {
    fields_ = "";
    orderBy_ = "";
    filter_ = "";
    view_ = "";
    q_ = "";
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private ListProjectMembersServiceRequest(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            fields_ = s;
            break;
          }
          case 18: {
            java.lang.String s = input.readStringRequireUtf8();

            orderBy_ = s;
            break;
          }
          case 26: {
            java.lang.String s = input.readStringRequireUtf8();

            filter_ = s;
            break;
          }
          case 32: {

            page_ = input.readInt32();
            break;
          }
          case 40: {

            limit_ = input.readInt32();
            break;
          }
          case 66: {
            java.lang.String s = input.readStringRequireUtf8();

            view_ = s;
            break;
          }
          case 90: {
            java.lang.String s = input.readStringRequireUtf8();

            q_ = s;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return taskmanager.AnyProto.internal_static_taskmanager_ListProjectMembersServiceRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return taskmanager.AnyProto.internal_static_taskmanager_ListProjectMembersServiceRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            taskmanager.ListProjectMembersServiceRequest.class, taskmanager.ListProjectMembersServiceRequest.Builder.class);
  }

  public static final int FIELDS_FIELD_NUMBER = 1;
  private volatile java.lang.Object fields_;
  /**
   * <pre>
   *Partial representation, fields=id,name
   * </pre>
   *
   * <code>string fields = 1;</code>
   */
  public java.lang.String getFields() {
    java.lang.Object ref = fields_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      fields_ = s;
      return s;
    }
  }
  /**
   * <pre>
   *Partial representation, fields=id,name
   * </pre>
   *
   * <code>string fields = 1;</code>
   */
  public com.google.protobuf.ByteString
      getFieldsBytes() {
    java.lang.Object ref = fields_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      fields_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int ORDER_BY_FIELD_NUMBER = 2;
  private volatile java.lang.Object orderBy_;
  /**
   * <pre>
   **
   * Sort fields, comma separated list for the ordering
   * use **?filter=-display_name** with a dash to sort descending
   * use **?filter=display_name** to sort ascending
   * </pre>
   *
   * <code>string order_by = 2;</code>
   */
  public java.lang.String getOrderBy() {
    java.lang.Object ref = orderBy_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      orderBy_ = s;
      return s;
    }
  }
  /**
   * <pre>
   **
   * Sort fields, comma separated list for the ordering
   * use **?filter=-display_name** with a dash to sort descending
   * use **?filter=display_name** to sort ascending
   * </pre>
   *
   * <code>string order_by = 2;</code>
   */
  public com.google.protobuf.ByteString
      getOrderByBytes() {
    java.lang.Object ref = orderBy_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      orderBy_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int FILTER_FIELD_NUMBER = 3;
  private volatile java.lang.Object filter_;
  /**
   * <pre>
   *Filter
   * </pre>
   *
   * <code>string filter = 3;</code>
   */
  public java.lang.String getFilter() {
    java.lang.Object ref = filter_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      filter_ = s;
      return s;
    }
  }
  /**
   * <pre>
   *Filter
   * </pre>
   *
   * <code>string filter = 3;</code>
   */
  public com.google.protobuf.ByteString
      getFilterBytes() {
    java.lang.Object ref = filter_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      filter_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int PAGE_FIELD_NUMBER = 4;
  private int page_;
  /**
   * <pre>
   *Page number for paginated content. Tipp: follow the HATEOAS next, prev,...
   * </pre>
   *
   * <code>int32 page = 4;</code>
   */
  public int getPage() {
    return page_;
  }

  public static final int LIMIT_FIELD_NUMBER = 5;
  private int limit_;
  /**
   * <pre>
   *Number of elements to return per page
   * </pre>
   *
   * <code>int32 limit = 5;</code>
   */
  public int getLimit() {
    return limit_;
  }

  public static final int VIEW_FIELD_NUMBER = 8;
  private volatile java.lang.Object view_;
  /**
   * <pre>
   *https://cloud.google.com/apis/design/design_patterns#resource_view
   * </pre>
   *
   * <code>string view = 8;</code>
   */
  public java.lang.String getView() {
    java.lang.Object ref = view_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      view_ = s;
      return s;
    }
  }
  /**
   * <pre>
   *https://cloud.google.com/apis/design/design_patterns#resource_view
   * </pre>
   *
   * <code>string view = 8;</code>
   */
  public com.google.protobuf.ByteString
      getViewBytes() {
    java.lang.Object ref = view_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      view_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int Q_FIELD_NUMBER = 11;
  private volatile java.lang.Object q_;
  /**
   * <pre>
   *Query term to search a member
   * </pre>
   *
   * <code>string q = 11;</code>
   */
  public java.lang.String getQ() {
    java.lang.Object ref = q_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      q_ = s;
      return s;
    }
  }
  /**
   * <pre>
   *Query term to search a member
   * </pre>
   *
   * <code>string q = 11;</code>
   */
  public com.google.protobuf.ByteString
      getQBytes() {
    java.lang.Object ref = q_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      q_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!getFieldsBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, fields_);
    }
    if (!getOrderByBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, orderBy_);
    }
    if (!getFilterBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, filter_);
    }
    if (page_ != 0) {
      output.writeInt32(4, page_);
    }
    if (limit_ != 0) {
      output.writeInt32(5, limit_);
    }
    if (!getViewBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 8, view_);
    }
    if (!getQBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 11, q_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getFieldsBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, fields_);
    }
    if (!getOrderByBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, orderBy_);
    }
    if (!getFilterBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, filter_);
    }
    if (page_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(4, page_);
    }
    if (limit_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(5, limit_);
    }
    if (!getViewBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(8, view_);
    }
    if (!getQBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(11, q_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof taskmanager.ListProjectMembersServiceRequest)) {
      return super.equals(obj);
    }
    taskmanager.ListProjectMembersServiceRequest other = (taskmanager.ListProjectMembersServiceRequest) obj;

    if (!getFields()
        .equals(other.getFields())) return false;
    if (!getOrderBy()
        .equals(other.getOrderBy())) return false;
    if (!getFilter()
        .equals(other.getFilter())) return false;
    if (getPage()
        != other.getPage()) return false;
    if (getLimit()
        != other.getLimit()) return false;
    if (!getView()
        .equals(other.getView())) return false;
    if (!getQ()
        .equals(other.getQ())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + FIELDS_FIELD_NUMBER;
    hash = (53 * hash) + getFields().hashCode();
    hash = (37 * hash) + ORDER_BY_FIELD_NUMBER;
    hash = (53 * hash) + getOrderBy().hashCode();
    hash = (37 * hash) + FILTER_FIELD_NUMBER;
    hash = (53 * hash) + getFilter().hashCode();
    hash = (37 * hash) + PAGE_FIELD_NUMBER;
    hash = (53 * hash) + getPage();
    hash = (37 * hash) + LIMIT_FIELD_NUMBER;
    hash = (53 * hash) + getLimit();
    hash = (37 * hash) + VIEW_FIELD_NUMBER;
    hash = (53 * hash) + getView().hashCode();
    hash = (37 * hash) + Q_FIELD_NUMBER;
    hash = (53 * hash) + getQ().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static taskmanager.ListProjectMembersServiceRequest parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(taskmanager.ListProjectMembersServiceRequest prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code taskmanager.ListProjectMembersServiceRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:taskmanager.ListProjectMembersServiceRequest)
      taskmanager.ListProjectMembersServiceRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return taskmanager.AnyProto.internal_static_taskmanager_ListProjectMembersServiceRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return taskmanager.AnyProto.internal_static_taskmanager_ListProjectMembersServiceRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              taskmanager.ListProjectMembersServiceRequest.class, taskmanager.ListProjectMembersServiceRequest.Builder.class);
    }

    // Construct using taskmanager.ListProjectMembersServiceRequest.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      fields_ = "";

      orderBy_ = "";

      filter_ = "";

      page_ = 0;

      limit_ = 0;

      view_ = "";

      q_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return taskmanager.AnyProto.internal_static_taskmanager_ListProjectMembersServiceRequest_descriptor;
    }

    @java.lang.Override
    public taskmanager.ListProjectMembersServiceRequest getDefaultInstanceForType() {
      return taskmanager.ListProjectMembersServiceRequest.getDefaultInstance();
    }

    @java.lang.Override
    public taskmanager.ListProjectMembersServiceRequest build() {
      taskmanager.ListProjectMembersServiceRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public taskmanager.ListProjectMembersServiceRequest buildPartial() {
      taskmanager.ListProjectMembersServiceRequest result = new taskmanager.ListProjectMembersServiceRequest(this);
      result.fields_ = fields_;
      result.orderBy_ = orderBy_;
      result.filter_ = filter_;
      result.page_ = page_;
      result.limit_ = limit_;
      result.view_ = view_;
      result.q_ = q_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof taskmanager.ListProjectMembersServiceRequest) {
        return mergeFrom((taskmanager.ListProjectMembersServiceRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(taskmanager.ListProjectMembersServiceRequest other) {
      if (other == taskmanager.ListProjectMembersServiceRequest.getDefaultInstance()) return this;
      if (!other.getFields().isEmpty()) {
        fields_ = other.fields_;
        onChanged();
      }
      if (!other.getOrderBy().isEmpty()) {
        orderBy_ = other.orderBy_;
        onChanged();
      }
      if (!other.getFilter().isEmpty()) {
        filter_ = other.filter_;
        onChanged();
      }
      if (other.getPage() != 0) {
        setPage(other.getPage());
      }
      if (other.getLimit() != 0) {
        setLimit(other.getLimit());
      }
      if (!other.getView().isEmpty()) {
        view_ = other.view_;
        onChanged();
      }
      if (!other.getQ().isEmpty()) {
        q_ = other.q_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      taskmanager.ListProjectMembersServiceRequest parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (taskmanager.ListProjectMembersServiceRequest) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private java.lang.Object fields_ = "";
    /**
     * <pre>
     *Partial representation, fields=id,name
     * </pre>
     *
     * <code>string fields = 1;</code>
     */
    public java.lang.String getFields() {
      java.lang.Object ref = fields_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        fields_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     *Partial representation, fields=id,name
     * </pre>
     *
     * <code>string fields = 1;</code>
     */
    public com.google.protobuf.ByteString
        getFieldsBytes() {
      java.lang.Object ref = fields_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        fields_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     *Partial representation, fields=id,name
     * </pre>
     *
     * <code>string fields = 1;</code>
     */
    public Builder setFields(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      fields_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Partial representation, fields=id,name
     * </pre>
     *
     * <code>string fields = 1;</code>
     */
    public Builder clearFields() {
      
      fields_ = getDefaultInstance().getFields();
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Partial representation, fields=id,name
     * </pre>
     *
     * <code>string fields = 1;</code>
     */
    public Builder setFieldsBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      fields_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object orderBy_ = "";
    /**
     * <pre>
     **
     * Sort fields, comma separated list for the ordering
     * use **?filter=-display_name** with a dash to sort descending
     * use **?filter=display_name** to sort ascending
     * </pre>
     *
     * <code>string order_by = 2;</code>
     */
    public java.lang.String getOrderBy() {
      java.lang.Object ref = orderBy_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        orderBy_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     **
     * Sort fields, comma separated list for the ordering
     * use **?filter=-display_name** with a dash to sort descending
     * use **?filter=display_name** to sort ascending
     * </pre>
     *
     * <code>string order_by = 2;</code>
     */
    public com.google.protobuf.ByteString
        getOrderByBytes() {
      java.lang.Object ref = orderBy_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        orderBy_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     **
     * Sort fields, comma separated list for the ordering
     * use **?filter=-display_name** with a dash to sort descending
     * use **?filter=display_name** to sort ascending
     * </pre>
     *
     * <code>string order_by = 2;</code>
     */
    public Builder setOrderBy(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      orderBy_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     **
     * Sort fields, comma separated list for the ordering
     * use **?filter=-display_name** with a dash to sort descending
     * use **?filter=display_name** to sort ascending
     * </pre>
     *
     * <code>string order_by = 2;</code>
     */
    public Builder clearOrderBy() {
      
      orderBy_ = getDefaultInstance().getOrderBy();
      onChanged();
      return this;
    }
    /**
     * <pre>
     **
     * Sort fields, comma separated list for the ordering
     * use **?filter=-display_name** with a dash to sort descending
     * use **?filter=display_name** to sort ascending
     * </pre>
     *
     * <code>string order_by = 2;</code>
     */
    public Builder setOrderByBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      orderBy_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object filter_ = "";
    /**
     * <pre>
     *Filter
     * </pre>
     *
     * <code>string filter = 3;</code>
     */
    public java.lang.String getFilter() {
      java.lang.Object ref = filter_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        filter_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     *Filter
     * </pre>
     *
     * <code>string filter = 3;</code>
     */
    public com.google.protobuf.ByteString
        getFilterBytes() {
      java.lang.Object ref = filter_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        filter_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     *Filter
     * </pre>
     *
     * <code>string filter = 3;</code>
     */
    public Builder setFilter(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      filter_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Filter
     * </pre>
     *
     * <code>string filter = 3;</code>
     */
    public Builder clearFilter() {
      
      filter_ = getDefaultInstance().getFilter();
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Filter
     * </pre>
     *
     * <code>string filter = 3;</code>
     */
    public Builder setFilterBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      filter_ = value;
      onChanged();
      return this;
    }

    private int page_ ;
    /**
     * <pre>
     *Page number for paginated content. Tipp: follow the HATEOAS next, prev,...
     * </pre>
     *
     * <code>int32 page = 4;</code>
     */
    public int getPage() {
      return page_;
    }
    /**
     * <pre>
     *Page number for paginated content. Tipp: follow the HATEOAS next, prev,...
     * </pre>
     *
     * <code>int32 page = 4;</code>
     */
    public Builder setPage(int value) {
      
      page_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Page number for paginated content. Tipp: follow the HATEOAS next, prev,...
     * </pre>
     *
     * <code>int32 page = 4;</code>
     */
    public Builder clearPage() {
      
      page_ = 0;
      onChanged();
      return this;
    }

    private int limit_ ;
    /**
     * <pre>
     *Number of elements to return per page
     * </pre>
     *
     * <code>int32 limit = 5;</code>
     */
    public int getLimit() {
      return limit_;
    }
    /**
     * <pre>
     *Number of elements to return per page
     * </pre>
     *
     * <code>int32 limit = 5;</code>
     */
    public Builder setLimit(int value) {
      
      limit_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Number of elements to return per page
     * </pre>
     *
     * <code>int32 limit = 5;</code>
     */
    public Builder clearLimit() {
      
      limit_ = 0;
      onChanged();
      return this;
    }

    private java.lang.Object view_ = "";
    /**
     * <pre>
     *https://cloud.google.com/apis/design/design_patterns#resource_view
     * </pre>
     *
     * <code>string view = 8;</code>
     */
    public java.lang.String getView() {
      java.lang.Object ref = view_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        view_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     *https://cloud.google.com/apis/design/design_patterns#resource_view
     * </pre>
     *
     * <code>string view = 8;</code>
     */
    public com.google.protobuf.ByteString
        getViewBytes() {
      java.lang.Object ref = view_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        view_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     *https://cloud.google.com/apis/design/design_patterns#resource_view
     * </pre>
     *
     * <code>string view = 8;</code>
     */
    public Builder setView(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      view_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *https://cloud.google.com/apis/design/design_patterns#resource_view
     * </pre>
     *
     * <code>string view = 8;</code>
     */
    public Builder clearView() {
      
      view_ = getDefaultInstance().getView();
      onChanged();
      return this;
    }
    /**
     * <pre>
     *https://cloud.google.com/apis/design/design_patterns#resource_view
     * </pre>
     *
     * <code>string view = 8;</code>
     */
    public Builder setViewBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      view_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object q_ = "";
    /**
     * <pre>
     *Query term to search a member
     * </pre>
     *
     * <code>string q = 11;</code>
     */
    public java.lang.String getQ() {
      java.lang.Object ref = q_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        q_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     *Query term to search a member
     * </pre>
     *
     * <code>string q = 11;</code>
     */
    public com.google.protobuf.ByteString
        getQBytes() {
      java.lang.Object ref = q_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        q_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     *Query term to search a member
     * </pre>
     *
     * <code>string q = 11;</code>
     */
    public Builder setQ(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      q_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Query term to search a member
     * </pre>
     *
     * <code>string q = 11;</code>
     */
    public Builder clearQ() {
      
      q_ = getDefaultInstance().getQ();
      onChanged();
      return this;
    }
    /**
     * <pre>
     *Query term to search a member
     * </pre>
     *
     * <code>string q = 11;</code>
     */
    public Builder setQBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      q_ = value;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:taskmanager.ListProjectMembersServiceRequest)
  }

  // @@protoc_insertion_point(class_scope:taskmanager.ListProjectMembersServiceRequest)
  private static final taskmanager.ListProjectMembersServiceRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new taskmanager.ListProjectMembersServiceRequest();
  }

  public static taskmanager.ListProjectMembersServiceRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ListProjectMembersServiceRequest>
      PARSER = new com.google.protobuf.AbstractParser<ListProjectMembersServiceRequest>() {
    @java.lang.Override
    public ListProjectMembersServiceRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new ListProjectMembersServiceRequest(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<ListProjectMembersServiceRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<ListProjectMembersServiceRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public taskmanager.ListProjectMembersServiceRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

