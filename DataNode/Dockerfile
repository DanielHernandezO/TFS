# Use an official Python runtime as a parent image
FROM python:3.9

# Set the working directory inside the container
WORKDIR /app

# Copy the requirements file into the container
COPY requirements.txt .

# Install the Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the protobuf files into the container
COPY protobufs/ protobufs/

ENV PYTHONUNBUFFERED 1
ENV MASTER_NAMENODE_URI 54.159.111.11:8081
ENV SLAVE_NAMENODE_URI  54.87.187.108:8081

RUN mkdir chunks/


# Compile the protobuf files
RUN python -m grpc_tools.protoc -I protobufs/  --python_out=. --grpc_python_out=. protobufs/datanode.proto
RUN python -m grpc_tools.protoc -I protobufs/ --python_out=. --grpc_python_out=. protobufs/namenode.proto
RUN ls
# Copy the main.py file into the container
COPY *.py /app/

# Run the Python application

CMD [ "python", "main.py" ]