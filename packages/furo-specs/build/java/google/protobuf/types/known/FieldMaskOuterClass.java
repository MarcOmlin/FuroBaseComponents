// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: google/protobuf/types/known/field_mask.proto

package google.protobuf.types.known;

public final class FieldMaskOuterClass {
  private FieldMaskOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface FieldMaskOrBuilder extends
      // @@protoc_insertion_point(interface_extends:google.protobuf.types.known.FieldMask)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    java.util.List<java.lang.String>
        getPathsList();
    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    int getPathsCount();
    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    java.lang.String getPaths(int index);
    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    com.google.protobuf.ByteString
        getPathsBytes(int index);
  }
  /**
   * <pre>
   * A field mask in update operations specifies which fields of the targeted resource are going to be updated. https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
   * </pre>
   *
   * Protobuf type {@code google.protobuf.types.known.FieldMask}
   */
  public  static final class FieldMask extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:google.protobuf.types.known.FieldMask)
      FieldMaskOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use FieldMask.newBuilder() to construct.
    private FieldMask(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private FieldMask() {
      paths_ = com.google.protobuf.LazyStringArrayList.EMPTY;
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private FieldMask(
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
              if (!((mutable_bitField0_ & 0x00000001) != 0)) {
                paths_ = new com.google.protobuf.LazyStringArrayList();
                mutable_bitField0_ |= 0x00000001;
              }
              paths_.add(s);
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
        if (((mutable_bitField0_ & 0x00000001) != 0)) {
          paths_ = paths_.getUnmodifiableView();
        }
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return google.protobuf.types.known.FieldMaskOuterClass.internal_static_google_protobuf_types_known_FieldMask_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return google.protobuf.types.known.FieldMaskOuterClass.internal_static_google_protobuf_types_known_FieldMask_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              google.protobuf.types.known.FieldMaskOuterClass.FieldMask.class, google.protobuf.types.known.FieldMaskOuterClass.FieldMask.Builder.class);
    }

    public static final int PATHS_FIELD_NUMBER = 1;
    private com.google.protobuf.LazyStringList paths_;
    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    public com.google.protobuf.ProtocolStringList
        getPathsList() {
      return paths_;
    }
    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    public int getPathsCount() {
      return paths_.size();
    }
    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    public java.lang.String getPaths(int index) {
      return paths_.get(index);
    }
    /**
     * <pre>
     * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
     * </pre>
     *
     * <code>repeated string paths = 1;</code>
     */
    public com.google.protobuf.ByteString
        getPathsBytes(int index) {
      return paths_.getByteString(index);
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
      for (int i = 0; i < paths_.size(); i++) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, paths_.getRaw(i));
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      {
        int dataSize = 0;
        for (int i = 0; i < paths_.size(); i++) {
          dataSize += computeStringSizeNoTag(paths_.getRaw(i));
        }
        size += dataSize;
        size += 1 * getPathsList().size();
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
      if (!(obj instanceof google.protobuf.types.known.FieldMaskOuterClass.FieldMask)) {
        return super.equals(obj);
      }
      google.protobuf.types.known.FieldMaskOuterClass.FieldMask other = (google.protobuf.types.known.FieldMaskOuterClass.FieldMask) obj;

      if (!getPathsList()
          .equals(other.getPathsList())) return false;
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
      if (getPathsCount() > 0) {
        hash = (37 * hash) + PATHS_FIELD_NUMBER;
        hash = (53 * hash) + getPathsList().hashCode();
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask parseFrom(
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
    public static Builder newBuilder(google.protobuf.types.known.FieldMaskOuterClass.FieldMask prototype) {
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
     * <pre>
     * A field mask in update operations specifies which fields of the targeted resource are going to be updated. https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
     * </pre>
     *
     * Protobuf type {@code google.protobuf.types.known.FieldMask}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:google.protobuf.types.known.FieldMask)
        google.protobuf.types.known.FieldMaskOuterClass.FieldMaskOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return google.protobuf.types.known.FieldMaskOuterClass.internal_static_google_protobuf_types_known_FieldMask_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return google.protobuf.types.known.FieldMaskOuterClass.internal_static_google_protobuf_types_known_FieldMask_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                google.protobuf.types.known.FieldMaskOuterClass.FieldMask.class, google.protobuf.types.known.FieldMaskOuterClass.FieldMask.Builder.class);
      }

      // Construct using google.protobuf.types.known.FieldMaskOuterClass.FieldMask.newBuilder()
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
        paths_ = com.google.protobuf.LazyStringArrayList.EMPTY;
        bitField0_ = (bitField0_ & ~0x00000001);
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return google.protobuf.types.known.FieldMaskOuterClass.internal_static_google_protobuf_types_known_FieldMask_descriptor;
      }

      @java.lang.Override
      public google.protobuf.types.known.FieldMaskOuterClass.FieldMask getDefaultInstanceForType() {
        return google.protobuf.types.known.FieldMaskOuterClass.FieldMask.getDefaultInstance();
      }

      @java.lang.Override
      public google.protobuf.types.known.FieldMaskOuterClass.FieldMask build() {
        google.protobuf.types.known.FieldMaskOuterClass.FieldMask result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public google.protobuf.types.known.FieldMaskOuterClass.FieldMask buildPartial() {
        google.protobuf.types.known.FieldMaskOuterClass.FieldMask result = new google.protobuf.types.known.FieldMaskOuterClass.FieldMask(this);
        int from_bitField0_ = bitField0_;
        if (((bitField0_ & 0x00000001) != 0)) {
          paths_ = paths_.getUnmodifiableView();
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.paths_ = paths_;
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
        if (other instanceof google.protobuf.types.known.FieldMaskOuterClass.FieldMask) {
          return mergeFrom((google.protobuf.types.known.FieldMaskOuterClass.FieldMask)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(google.protobuf.types.known.FieldMaskOuterClass.FieldMask other) {
        if (other == google.protobuf.types.known.FieldMaskOuterClass.FieldMask.getDefaultInstance()) return this;
        if (!other.paths_.isEmpty()) {
          if (paths_.isEmpty()) {
            paths_ = other.paths_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensurePathsIsMutable();
            paths_.addAll(other.paths_);
          }
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
        google.protobuf.types.known.FieldMaskOuterClass.FieldMask parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (google.protobuf.types.known.FieldMaskOuterClass.FieldMask) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }
      private int bitField0_;

      private com.google.protobuf.LazyStringList paths_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      private void ensurePathsIsMutable() {
        if (!((bitField0_ & 0x00000001) != 0)) {
          paths_ = new com.google.protobuf.LazyStringArrayList(paths_);
          bitField0_ |= 0x00000001;
         }
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public com.google.protobuf.ProtocolStringList
          getPathsList() {
        return paths_.getUnmodifiableView();
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public int getPathsCount() {
        return paths_.size();
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public java.lang.String getPaths(int index) {
        return paths_.get(index);
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public com.google.protobuf.ByteString
          getPathsBytes(int index) {
        return paths_.getByteString(index);
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public Builder setPaths(
          int index, java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  ensurePathsIsMutable();
        paths_.set(index, value);
        onChanged();
        return this;
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public Builder addPaths(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  ensurePathsIsMutable();
        paths_.add(value);
        onChanged();
        return this;
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public Builder addAllPaths(
          java.lang.Iterable<java.lang.String> values) {
        ensurePathsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, paths_);
        onChanged();
        return this;
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public Builder clearPaths() {
        paths_ = com.google.protobuf.LazyStringArrayList.EMPTY;
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
        return this;
      }
      /**
       * <pre>
       * The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
       * </pre>
       *
       * <code>repeated string paths = 1;</code>
       */
      public Builder addPathsBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        ensurePathsIsMutable();
        paths_.add(value);
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


      // @@protoc_insertion_point(builder_scope:google.protobuf.types.known.FieldMask)
    }

    // @@protoc_insertion_point(class_scope:google.protobuf.types.known.FieldMask)
    private static final google.protobuf.types.known.FieldMaskOuterClass.FieldMask DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new google.protobuf.types.known.FieldMaskOuterClass.FieldMask();
    }

    public static google.protobuf.types.known.FieldMaskOuterClass.FieldMask getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<FieldMask>
        PARSER = new com.google.protobuf.AbstractParser<FieldMask>() {
      @java.lang.Override
      public FieldMask parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new FieldMask(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<FieldMask> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<FieldMask> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public google.protobuf.types.known.FieldMaskOuterClass.FieldMask getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_google_protobuf_types_known_FieldMask_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_google_protobuf_types_known_FieldMask_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n,google/protobuf/types/known/field_mask" +
      ".proto\022\033google.protobuf.types.known\"\032\n\tF" +
      "ieldMask\022\r\n\005paths\030\001 \003(\tb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        }, assigner);
    internal_static_google_protobuf_types_known_FieldMask_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_google_protobuf_types_known_FieldMask_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_google_protobuf_types_known_FieldMask_descriptor,
        new java.lang.String[] { "Paths", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
